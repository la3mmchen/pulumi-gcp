// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
func LookupBucketObject(ctx *pulumi.Context, args *GetBucketObjectArgs) (*GetBucketObjectResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["bucket"] = args.Bucket
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("gcp:storage/getBucketObject:getBucketObject", inputs)
	if err != nil {
		return nil, err
	}
	return &GetBucketObjectResult{
		Bucket: outputs["bucket"],
		CacheControl: outputs["cacheControl"],
		Content: outputs["content"],
		ContentDisposition: outputs["contentDisposition"],
		ContentEncoding: outputs["contentEncoding"],
		ContentLanguage: outputs["contentLanguage"],
		ContentType: outputs["contentType"],
		Crc32c: outputs["crc32c"],
		DetectMd5hash: outputs["detectMd5hash"],
		Md5hash: outputs["md5hash"],
		Name: outputs["name"],
		OutputName: outputs["outputName"],
		PredefinedAcl: outputs["predefinedAcl"],
		SelfLink: outputs["selfLink"],
		Source: outputs["source"],
		StorageClass: outputs["storageClass"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getBucketObject.
type GetBucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket interface{}
	// The name of the object.
	Name interface{}
}

// A collection of values returned by getBucketObject.
type GetBucketObjectResult struct {
	Bucket interface{}
	// (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl interface{}
	Content interface{}
	// (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition interface{}
	// (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding interface{}
	// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage interface{}
	// (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType interface{}
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c interface{}
	DetectMd5hash interface{}
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash interface{}
	Name interface{}
	OutputName interface{}
	PredefinedAcl interface{}
	// (Computed) A url reference to this object.
	SelfLink interface{}
	Source interface{}
	// (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}