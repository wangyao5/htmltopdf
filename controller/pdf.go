package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
)

// DownloadPdf 通过html url地址下载pdf附件
func DownloadPdf(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "url is empty",
		})
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	// pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage(url)

	// Set options for this page
	// page.FooterRight.Set("[page]")
	// page.FooterFontSize.Set(10)
	// page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Accept-Ranges", "bytes")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s", "download.pdf")) //文件名
	c.Writer.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Write(pdfg.Bytes())
}
