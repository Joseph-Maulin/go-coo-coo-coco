package parser

/*
	{
	  "info": {...},
	  "licenses": [...],
	  "images": [...],
	  "annotations": [...],
	  "categories": [...]
	}
*/
type CocoJson struct {
	Info        CocoJsonInfo         `json:"info"`        // Information about the dataset
	Licenses    []CocoJsonLicense    `json:"licenses"`    // Licenses for the dataset
	Images      []CocoJsonImage      `json:"images"`      // Images in the dataset
	Annotations []CocoJsonAnnotation `json:"annotations"` // Annotations for the dataset
	Categories  []CocoJsonCategory   `json:"categories"`  // Categories for the dataset
}

/*
	"info": {
	  "description": "Example Dataset",
	  "version": "1.0",
	  "year": 2025,
	  "contributor": "Alex Dev",
	  "date_created": "2025-11-05"
	}
*/
type CocoJsonInfo struct {
	Year        int    `json:"year"`         // Year of the dataset
	Version     string `json:"version"`      // Version of the dataset
	Description string `json:"description"`  // Description of the dataset
	Contributor string `json:"contributor"`  // Contributor of the dataset
	URL         string `json:"url"`          // URL of the dataset
	DateCreated string `json:"date_created"` // ISO date string
}

/*
"licenses": [

	{
	  "id": 1,
	  "name": "CC-BY-4.0",
	  "url": "https://creativecommons.org/licenses/by/4.0/"
	}

]
*/
type CocoJsonLicense struct {
	ID   int    `json:"id"`   // License ID
	Name string `json:"name"` // License name
	URL  string `json:"url"`  // License URL
}

/*
"images": [

	{
	  "id": 1,
	  "width": 1280,
	  "height": 720,
	  "file_name": "image_001.jpg",
	  "license": 1,
	  "date_captured": "2025-11-05"
	}

]
*/
type CocoJsonImage struct {
	ID           int    `json:"id"`                   // Image ID
	Width        int    `json:"width"`                // Image width in pixels
	Height       int    `json:"height"`               // Image height in pixels
	FileName     string `json:"file_name"`            // Filename of the image or path to the image
	License      int    `json:"license"`              // License ID (matches licenses.id)
	FlickrURL    string `json:"flickr_url,omitempty"` // (Optional) Flickr URL for the image
	CocoURL      string `json:"coco_url,omitempty"`   // (Optional) COCO URL for the image
	DateCaptured string `json:"date_captured"`        // ISO date string of when the image was captured
}

/*
"categories": [

	{
	  "id": 1,
	  "name": "person",
	  "supercategory": "human",
	  "keypoints": ["nose", "left_eye", "right_eye", "left_ear", "right_ear"],
	  "skeleton": [[1,2],[1,3],[2,4],[3,5]]
	}

]
*/
type CocoJsonCategory struct {
	ID            int      `json:"id"`                  // Category ID
	Name          string   `json:"name"`                // Category name
	SuperCategory string   `json:"supercategory"`       // Supercategory name
	Keypoints     []string `json:"keypoints,omitempty"` // (Optional) For keypoint datasets
	Skeleton      [][]int  `json:"skeleton,omitempty"`  // (Optional) For keypoint connectivity
}

/*
		"annotations": [
	  {
	    "id": 202,
	    "image_id": 1,
	    "category_id": 1,
	    "keypoints": [100, 200, 2, 120, 210, 2, 130, 220, 1],
	    "num_keypoints": 3,
	    "bbox": [80, 180, 100, 120],
	    "area": 12000.0,
	    "iscrowd": 0
	  }

]
*/
type CocoJsonAnnotation struct {
	ID           int            `json:"id"`                      // Annotation ID
	ImageID      int            `json:"image_id"`                // ID of the image (matches images.id)
	CategoryID   int            `json:"category_id"`             // ID of the category (matches categories.id)
	Segmentation [][]float64    `json:"segmentation,omitempty"`  // Polygon or RLE mask
	BBox         []float64      `json:"bbox,omitempty"`          // Bounding box in pixels [x, y, width, height]
	Points       [][]float64    `json:"points,omitempty"`        // List of points for the box (e.g. [ [x1, y1], [x2, y2], ... ])
	Area         float64        `json:"area,omitempty"`          // Area of the object (in pixels²)
	IsCrowd      int            `json:"iscrowd,omitempty"`       // 0 = individual, 1 = crowd/merged
	Attributes   map[string]any `json:"attributes,omitempty"`    // (Optional) Attributes for the object (e.g. { "color": "red", "size": "large" })
	Keypoints    []int          `json:"keypoints,omitempty"`     // (Optional) For pose/keypoint datasets
	NumKeypoints int            `json:"num_keypoints,omitempty"` // (Optional) Number of labeled keypoints
}
