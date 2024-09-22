package helpers

import (
	"context"
	"fmt"
	"log"
	"net/url"

	vision "cloud.google.com/go/vision/apiv1"
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
)

func getBaseURL(inputURL string) string {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	return baseURL
}

func PrintResults(webAnnotations *pb.WebDetection) ([]string, []string) {
	var fullMatchingLinks []string
	var imageMatchingLinks []string
	fmt.Println("Pages with matching images:")
	for _, page := range webAnnotations.PagesWithMatchingImages {
		fmt.Printf("\tURL: %s\n", page.Url)
		fullMatchingLinks = append(fullMatchingLinks, page.Url)
	}

	fmt.Println("\nFull matches found:")
	for _, image := range webAnnotations.FullMatchingImages {
		fmt.Printf("\tURL: %s\n", image.Url)
		baseURL := getBaseURL(image.Url)
		imageMatchingLinks = append(imageMatchingLinks, baseURL)
	}

	return fullMatchingLinks, imageMatchingLinks
}

func DetectWeb(ctx context.Context, c *vision.ImageAnnotatorClient, imageURI string) *pb.WebDetection {
	image := &pb.Image{Source: &pb.ImageSource{ImageUri: imageURI}}
	features := []*pb.Feature{
		{Type: pb.Feature_WEB_DETECTION},
	}
	req := &pb.AnnotateImageRequest{
		Image:        image,
		Features:     features,
		ImageContext: &pb.ImageContext{},
	}

	resp, err := c.AnnotateImage(ctx, req)
	if err != nil {
		log.Fatalf("Failed to annotate image: %v", err)
	}

	return resp.WebDetection
}