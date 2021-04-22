// Code generated by go generate via internal/cmd/service; DO NOT EDIT.
package gcs

import (
	"context"
	"io"

	"github.com/aos-dev/go-storage/v3/pkg/credential"
	"github.com/aos-dev/go-storage/v3/pkg/endpoint"
	"github.com/aos-dev/go-storage/v3/pkg/httpclient"
	"github.com/aos-dev/go-storage/v3/services"
	. "github.com/aos-dev/go-storage/v3/types"
)

var _ credential.Provider
var _ endpoint.Value
var _ Storager
var _ services.ServiceError
var _ httpclient.Options

// Type is the type for gcs
const Type = "gcs"

// Service available pairs.
const (
	// DefaultServicePairs set default pairs for service actions
	pairDefaultServicePairs = "gcs_default_service_pairs"
	// DefaultStoragePairs set default pairs for storager actions
	pairDefaultStoragePairs = "gcs_default_storage_pairs"
	// EncryptionKey is the customer's 32-byte AES-256 key
	pairEncryptionKey = "gcs_encryption_key"
	// ProjectID
	pairProjectID = "gcs_project_id"
	// StorageClass
	pairStorageClass = "gcs_storage_class"
)

// Service available metadata.
const (
	MetadataEncryptionKeySha256 = "gcs-encryption_key_sha256"

	MetadataStorageClass = "gcs-storage-class"
)

// WithDefaultServicePairs will apply default_service_pairs value to Options
// DefaultServicePairs set default pairs for service actions
func WithDefaultServicePairs(v DefaultServicePairs) Pair {
	return Pair{
		Key:   pairDefaultServicePairs,
		Value: v,
	}
}

// WithDefaultStoragePairs will apply default_storage_pairs value to Options
// DefaultStoragePairs set default pairs for storager actions
func WithDefaultStoragePairs(v DefaultStoragePairs) Pair {
	return Pair{
		Key:   pairDefaultStoragePairs,
		Value: v,
	}
}

// WithEncryptionKey will apply encryption_key value to Options
// EncryptionKey is the customer's 32-byte AES-256 key
func WithEncryptionKey(v []byte) Pair {
	return Pair{
		Key:   pairEncryptionKey,
		Value: v,
	}
}

// WithProjectID will apply project_id value to Options
// ProjectID
func WithProjectID(v string) Pair {
	return Pair{
		Key:   pairProjectID,
		Value: v,
	}
}

// WithStorageClass will apply storage_class value to Options
// StorageClass
func WithStorageClass(v string) Pair {
	return Pair{
		Key:   pairStorageClass,
		Value: v,
	}
}

// pairServiceNew is the parsed struct
type pairServiceNew struct {
	pairs []Pair

	// Required pairs
	HasCredential bool
	Credential    string
	HasProjectID  bool
	ProjectID     string
	// Optional pairs
	HasDefaultServicePairs bool
	DefaultServicePairs    DefaultServicePairs
	HasHTTPClientOptions   bool
	HTTPClientOptions      *httpclient.Options
	// Generated pairs
}

// parsePairServiceNew will parse Pair slice into *pairServiceNew
func parsePairServiceNew(opts []Pair) (pairServiceNew, error) {
	result := pairServiceNew{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		case "credential":
			if result.HasCredential {
				continue
			}
			result.HasCredential = true
			result.Credential = v.Value.(string)
		case pairProjectID:
			if result.HasProjectID {
				continue
			}
			result.HasProjectID = true
			result.ProjectID = v.Value.(string)
		// Optional pairs
		case pairDefaultServicePairs:
			if result.HasDefaultServicePairs {
				continue
			}
			result.HasDefaultServicePairs = true
			result.DefaultServicePairs = v.Value.(DefaultServicePairs)
		case "http_client_options":
			if result.HasHTTPClientOptions {
				continue
			}
			result.HasHTTPClientOptions = true
			result.HTTPClientOptions = v.Value.(*httpclient.Options)
			// Generated pairs
		}
	}
	if !result.HasCredential {
		return pairServiceNew{}, services.NewPairRequiredError("credential")
	}
	if !result.HasProjectID {
		return pairServiceNew{}, services.NewPairRequiredError(pairProjectID)
	}

	return result, nil
}

// DefaultServicePairs is default pairs for specific action
type DefaultServicePairs struct {
	Create []Pair
	Delete []Pair
	Get    []Pair
	List   []Pair
}

// pairServiceCreate is the parsed struct
type pairServiceCreate struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairServiceCreate will parse Pair slice into *pairServiceCreate
func (s *Service) parsePairServiceCreate(opts []Pair) (pairServiceCreate, error) {
	result := pairServiceCreate{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			continue

		}
	}

	return result, nil
}

// pairServiceDelete is the parsed struct
type pairServiceDelete struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairServiceDelete will parse Pair slice into *pairServiceDelete
func (s *Service) parsePairServiceDelete(opts []Pair) (pairServiceDelete, error) {
	result := pairServiceDelete{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			continue

		}
	}

	return result, nil
}

// pairServiceGet is the parsed struct
type pairServiceGet struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairServiceGet will parse Pair slice into *pairServiceGet
func (s *Service) parsePairServiceGet(opts []Pair) (pairServiceGet, error) {
	result := pairServiceGet{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			continue

		}
	}

	return result, nil
}

// pairServiceList is the parsed struct
type pairServiceList struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairServiceList will parse Pair slice into *pairServiceList
func (s *Service) parsePairServiceList(opts []Pair) (pairServiceList, error) {
	result := pairServiceList{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			continue

		}
	}

	return result, nil
}

// Create will create a new storager instance.
//
// This function will create a context by default.
func (s *Service) Create(name string, pairs ...Pair) (store Storager, err error) {
	ctx := context.Background()
	return s.CreateWithContext(ctx, name, pairs...)
}

// CreateWithContext will create a new storager instance.
func (s *Service) CreateWithContext(ctx context.Context, name string, pairs ...Pair) (store Storager, err error) {
	pairs = append(pairs, s.defaultPairs.Create...)
	var opt pairServiceCreate

	defer func() {
		err = s.formatError("create", err, name)
	}()

	opt, err = s.parsePairServiceCreate(pairs)
	if err != nil {
		return
	}

	return s.create(ctx, name, opt)
}

// Delete will delete a storager instance.
//
// This function will create a context by default.
func (s *Service) Delete(name string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, name, pairs...)
}

// DeleteWithContext will delete a storager instance.
func (s *Service) DeleteWithContext(ctx context.Context, name string, pairs ...Pair) (err error) {
	pairs = append(pairs, s.defaultPairs.Delete...)
	var opt pairServiceDelete

	defer func() {
		err = s.formatError("delete", err, name)
	}()

	opt, err = s.parsePairServiceDelete(pairs)
	if err != nil {
		return
	}

	return s.delete(ctx, name, opt)
}

// Get will get a valid storager instance for service.
//
// This function will create a context by default.
func (s *Service) Get(name string, pairs ...Pair) (store Storager, err error) {
	ctx := context.Background()
	return s.GetWithContext(ctx, name, pairs...)
}

// GetWithContext will get a valid storager instance for service.
func (s *Service) GetWithContext(ctx context.Context, name string, pairs ...Pair) (store Storager, err error) {
	pairs = append(pairs, s.defaultPairs.Get...)
	var opt pairServiceGet

	defer func() {
		err = s.formatError("get", err, name)
	}()

	opt, err = s.parsePairServiceGet(pairs)
	if err != nil {
		return
	}

	return s.get(ctx, name, opt)
}

// List will list all storager instances under this service.
//
// This function will create a context by default.
func (s *Service) List(pairs ...Pair) (sti *StoragerIterator, err error) {
	ctx := context.Background()
	return s.ListWithContext(ctx, pairs...)
}

// ListWithContext will list all storager instances under this service.
func (s *Service) ListWithContext(ctx context.Context, pairs ...Pair) (sti *StoragerIterator, err error) {
	pairs = append(pairs, s.defaultPairs.List...)
	var opt pairServiceList

	defer func() {

		err = s.formatError("list", err, "")
	}()

	opt, err = s.parsePairServiceList(pairs)
	if err != nil {
		return
	}

	return s.list(ctx, opt)
}

// pairStorageNew is the parsed struct
type pairStorageNew struct {
	pairs []Pair

	// Required pairs
	HasName bool
	Name    string
	// Optional pairs
	HasDefaultStoragePairs bool
	DefaultStoragePairs    DefaultStoragePairs
	HasPairPolicy          bool
	PairPolicy             PairPolicy
	HasWorkDir             bool
	WorkDir                string
	// Generated pairs
}

// parsePairStorageNew will parse Pair slice into *pairStorageNew
func parsePairStorageNew(opts []Pair) (pairStorageNew, error) {
	result := pairStorageNew{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		case "name":
			if result.HasName {
				continue
			}
			result.HasName = true
			result.Name = v.Value.(string)
		// Optional pairs
		case pairDefaultStoragePairs:
			if result.HasDefaultStoragePairs {
				continue
			}
			result.HasDefaultStoragePairs = true
			result.DefaultStoragePairs = v.Value.(DefaultStoragePairs)
		case "pair_policy":
			if result.HasPairPolicy {
				continue
			}
			result.HasPairPolicy = true
			result.PairPolicy = v.Value.(PairPolicy)
		case "work_dir":
			if result.HasWorkDir {
				continue
			}
			result.HasWorkDir = true
			result.WorkDir = v.Value.(string)
			// Generated pairs
		}
	}
	if !result.HasName {
		return pairStorageNew{}, services.NewPairRequiredError("name")
	}

	return result, nil
}

// DefaultStoragePairs is default pairs for specific action
type DefaultStoragePairs struct {
	Create   []Pair
	Delete   []Pair
	List     []Pair
	Metadata []Pair
	Read     []Pair
	Stat     []Pair
	Write    []Pair
}

// pairStorageCreate is the parsed struct
type pairStorageCreate struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageCreate will parse Pair slice into *pairStorageCreate
func (s *Storage) parsePairStorageCreate(opts []Pair) (pairStorageCreate, error) {
	result := pairStorageCreate{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			if s.pairPolicy.All || s.pairPolicy.Create {
				return pairStorageCreate{}, services.NewPairUnsupportedError(v)
			}

		}
	}

	return result, nil
}

// pairStorageDelete is the parsed struct
type pairStorageDelete struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageDelete will parse Pair slice into *pairStorageDelete
func (s *Storage) parsePairStorageDelete(opts []Pair) (pairStorageDelete, error) {
	result := pairStorageDelete{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			if s.pairPolicy.All || s.pairPolicy.Delete {
				return pairStorageDelete{}, services.NewPairUnsupportedError(v)
			}

		}
	}

	return result, nil
}

// pairStorageList is the parsed struct
type pairStorageList struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	HasListMode bool
	ListMode    ListMode
	// Generated pairs
}

// parsePairStorageList will parse Pair slice into *pairStorageList
func (s *Storage) parsePairStorageList(opts []Pair) (pairStorageList, error) {
	result := pairStorageList{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		case "list_mode":
			result.HasListMode = true
			result.ListMode = v.Value.(ListMode)
		// Generated pairs
		default:

			if s.pairPolicy.All || s.pairPolicy.List {
				return pairStorageList{}, services.NewPairUnsupportedError(v)
			}

		}
	}

	return result, nil
}

// pairStorageMetadata is the parsed struct
type pairStorageMetadata struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageMetadata will parse Pair slice into *pairStorageMetadata
func (s *Storage) parsePairStorageMetadata(opts []Pair) (pairStorageMetadata, error) {
	result := pairStorageMetadata{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			if s.pairPolicy.All || s.pairPolicy.Metadata {
				return pairStorageMetadata{}, services.NewPairUnsupportedError(v)
			}

		}
	}

	return result, nil
}

// pairStorageRead is the parsed struct
type pairStorageRead struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	HasEncryptionKey bool
	EncryptionKey    []byte
	HasIoCallback    bool
	IoCallback       func([]byte)
	HasOffset        bool
	Offset           int64
	HasSize          bool
	Size             int64
	// Generated pairs
}

// parsePairStorageRead will parse Pair slice into *pairStorageRead
func (s *Storage) parsePairStorageRead(opts []Pair) (pairStorageRead, error) {
	result := pairStorageRead{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		case pairEncryptionKey:
			result.HasEncryptionKey = true
			result.EncryptionKey = v.Value.([]byte)
		case "io_callback":
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
		case "offset":
			result.HasOffset = true
			result.Offset = v.Value.(int64)
		case "size":
			result.HasSize = true
			result.Size = v.Value.(int64)
		// Generated pairs
		default:

			if s.pairPolicy.All || s.pairPolicy.Read {
				return pairStorageRead{}, services.NewPairUnsupportedError(v)
			}

		}
	}

	return result, nil
}

// pairStorageStat is the parsed struct
type pairStorageStat struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageStat will parse Pair slice into *pairStorageStat
func (s *Storage) parsePairStorageStat(opts []Pair) (pairStorageStat, error) {
	result := pairStorageStat{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		// Generated pairs
		default:

			if s.pairPolicy.All || s.pairPolicy.Stat {
				return pairStorageStat{}, services.NewPairUnsupportedError(v)
			}

		}
	}

	return result, nil
}

// pairStorageWrite is the parsed struct
type pairStorageWrite struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	HasContentMd5    bool
	ContentMd5       string
	HasContentType   bool
	ContentType      string
	HasEncryptionKey bool
	EncryptionKey    []byte
	HasIoCallback    bool
	IoCallback       func([]byte)
	HasStorageClass  bool
	StorageClass     string
	// Generated pairs
}

// parsePairStorageWrite will parse Pair slice into *pairStorageWrite
func (s *Storage) parsePairStorageWrite(opts []Pair) (pairStorageWrite, error) {
	result := pairStorageWrite{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		// Optional pairs
		case "content_md5":
			result.HasContentMd5 = true
			result.ContentMd5 = v.Value.(string)
		case "content_type":
			result.HasContentType = true
			result.ContentType = v.Value.(string)
		case pairEncryptionKey:
			result.HasEncryptionKey = true
			result.EncryptionKey = v.Value.([]byte)
		case "io_callback":
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
		case pairStorageClass:
			result.HasStorageClass = true
			result.StorageClass = v.Value.(string)
		// Generated pairs
		default:

			if s.pairPolicy.All || s.pairPolicy.Write {
				return pairStorageWrite{}, services.NewPairUnsupportedError(v)
			}

		}
	}

	return result, nil
}

// Create will create a new object without any api call.
//
// This function will create a context by default.
func (s *Storage) Create(path string, pairs ...Pair) (o *Object) {
	pairs = append(pairs, s.defaultPairs.Create...)
	var opt pairStorageCreate

	// Ignore error while handling local funtions.
	opt, _ = s.parsePairStorageCreate(pairs)

	return s.create(path, opt)
}

// Delete will delete an Object from service.
//
// This function will create a context by default.
func (s *Storage) Delete(path string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, path, pairs...)
}

// DeleteWithContext will delete an Object from service.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...Pair) (err error) {
	pairs = append(pairs, s.defaultPairs.Delete...)
	var opt pairStorageDelete

	defer func() {
		err = s.formatError("delete", err, path)
	}()

	opt, err = s.parsePairStorageDelete(pairs)
	if err != nil {
		return
	}

	return s.delete(ctx, path, opt)
}

// List will return list a specific path.
//
// This function will create a context by default.
func (s *Storage) List(path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	ctx := context.Background()
	return s.ListWithContext(ctx, path, pairs...)
}

// ListWithContext will return list a specific path.
func (s *Storage) ListWithContext(ctx context.Context, path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	pairs = append(pairs, s.defaultPairs.List...)
	var opt pairStorageList

	defer func() {
		err = s.formatError("list", err, path)
	}()

	opt, err = s.parsePairStorageList(pairs)
	if err != nil {
		return
	}

	return s.list(ctx, path, opt)
}

// Metadata will return current storager metadata.
//
// This function will create a context by default.
func (s *Storage) Metadata(pairs ...Pair) (meta *StorageMeta, err error) {
	ctx := context.Background()
	return s.MetadataWithContext(ctx, pairs...)
}

// MetadataWithContext will return current storager metadata.
func (s *Storage) MetadataWithContext(ctx context.Context, pairs ...Pair) (meta *StorageMeta, err error) {
	pairs = append(pairs, s.defaultPairs.Metadata...)
	var opt pairStorageMetadata

	defer func() {
		err = s.formatError("metadata", err)
	}()

	opt, err = s.parsePairStorageMetadata(pairs)
	if err != nil {
		return
	}

	return s.metadata(ctx, opt)
}

// Read will read the file's data.
//
// This function will create a context by default.
func (s *Storage) Read(path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.ReadWithContext(ctx, path, w, pairs...)
}

// ReadWithContext will read the file's data.
func (s *Storage) ReadWithContext(ctx context.Context, path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	pairs = append(pairs, s.defaultPairs.Read...)
	var opt pairStorageRead

	defer func() {
		err = s.formatError("read", err, path)
	}()

	opt, err = s.parsePairStorageRead(pairs)
	if err != nil {
		return
	}

	return s.read(ctx, path, w, opt)
}

// Stat will stat a path to get info of an object.
//
// This function will create a context by default.
func (s *Storage) Stat(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.StatWithContext(ctx, path, pairs...)
}

// StatWithContext will stat a path to get info of an object.
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	pairs = append(pairs, s.defaultPairs.Stat...)
	var opt pairStorageStat

	defer func() {
		err = s.formatError("stat", err, path)
	}()

	opt, err = s.parsePairStorageStat(pairs)
	if err != nil {
		return
	}

	return s.stat(ctx, path, opt)
}

// Write will write data into a file.
//
// This function will create a context by default.
func (s *Storage) Write(path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.WriteWithContext(ctx, path, r, size, pairs...)
}

// WriteWithContext will write data into a file.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	pairs = append(pairs, s.defaultPairs.Write...)
	var opt pairStorageWrite

	defer func() {
		err = s.formatError("write", err, path)
	}()

	opt, err = s.parsePairStorageWrite(pairs)
	if err != nil {
		return
	}

	return s.write(ctx, path, r, size, opt)
}
