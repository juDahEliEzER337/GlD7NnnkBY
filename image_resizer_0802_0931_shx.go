// 代码生成时间: 2025-08-02 09:31:37
package main

import (
# 扩展功能模块
    "fmt"
    "image"
    "image/jpeg"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// ImageResizer is a struct that holds the configuration for resizing images.
type ImageResizer struct {
    OutputPath string
    Width      int
    Height     int
}

// NewImageResizer initializes a new ImageResizer with the given output path and dimensions.
func NewImageResizer(outputPath string, width, height int) *ImageResizer {
    return &ImageResizer{
        OutputPath: outputPath,
        Width:      width,
        Height:     height,
    }
}

// Resize resizes an image to the specified width and height.
func (ir *ImageResizer) Resize(imagePath string) error {
    // Open the image file.
    imgFile, err := os.Open(imagePath)
    if err != nil {
        return err
