// 代码生成时间: 2025-08-24 13:52:03
// decompress_tool.go
package main

import (
    "archive/tar"
    "archive/zip"
    "bufio"
    "compress/gzip"
    "io"
    "io/fs"
    "log"
    "os"
    "path"
    "path/filepath"
)

// Decompress decompresses a file based on its extension.
func Decompress(inputPath, outputPath string) error {
    // Determine the file extension and handle accordingly.
    switch path.Ext(inputPath) {
    case ".zip":
        return unzip(inputPath, outputPath)
    case ".tar.gz", ".tgz":
        return untargz(inputPath, outputPath)
    case ".tar":
        return untar(inputPath, outputPath)
    default:
        return fmt.Errorf("unsupported file type: %s", path.Ext(inputPath))
    }
}

// unzip decompresses a zip file.
func unzip(inputPath, outputPath string) error {
    file, err := zip.OpenReader(inputPath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    for _, f := range file.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        
        // Create the output directory structure.
        targetPath := filepath.Join(outputPath, f.Name)
        
        if f.FileInfo().IsDir() {
            os.MkdirAll(targetPath, os.ModePerm)
        } else {
            os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
            outFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer outFile.Close()
            
            _, err = io.Copy(outFile, rc)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// untargz decompresses a gzipped tar file.
func untargz(inputPath, outputPath string) error {
    f, err := os.Open(inputPath)
    if err != nil {
        return err
    }
    defer f.Close()
    
    gz, err := gzip.NewReader(f)
    if err != nil {
        return err
    }
    defer gz.Close()
    
    return untar(gz, outputPath)
}

// untar decompresses a tar file.
func untar(input io.Reader, outputPath string) error {
    tr := tar.NewReader(input)
    for {
        hdr, err := tr.Next()
        if err == io.EOF {
            break // End of archive
        }
        if err != nil {
            return err
        }
        
        target := filepath.Join(outputPath, hdr.Name)
        switch hdr.Typeflag {
        case tar.TypeDir:
            if err := os.MkdirAll(target, os.ModePerm); err != nil {
                return err
            }
        case tar.TypeReg:
            outFile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(hdr.Mode));
            if err != nil {
                return err
            }
            _, err = io.Copy(outFile, tr)
            outFile.Close()
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func main() {
    inputFile := "path/to/your/input/file"
    outputDir := "path/to/your/output/directory"
    err := Decompress(inputFile, outputDir)
    if err != nil {
        log.Fatal(err)
    }
}
