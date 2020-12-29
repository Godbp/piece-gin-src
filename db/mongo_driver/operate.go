package mongo_driver

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DefaultCtx 默认上下文环境生成
func (d *Dao) DefaultCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// DecodeOne 处理DecodeOne的错误
// deprecated 该方法容易引起panic，已经废弃
func (d *Dao) DecodeOne(err error) error {
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err == nil {
		return nil
	}
	return err
}

// HasDuplicatedError 是否是重复id写入错误
func (d *Dao) HasDuplicatedError(err error) bool {
	if err, ok := err.(mongo.WriteException); ok {
		for _, e := range err.WriteErrors {
			if e.Code == 11000 {
				return true
			}
		}
	}
	return false
}

// IsAllDuplicatedError 是否是重复id写入错误
func (d *Dao) IsAllDuplicatedError(err error) bool {
	switch err.(type) {
	case mongo.WriteException:
		for _, e := range err.(mongo.WriteException).WriteErrors {
			if e.Code != 11000 {
				return false
			}
		}
		return true
	case mongo.BulkWriteException:
		for _, e := range err.(mongo.BulkWriteException).WriteErrors {
			if e.Code != 11000 {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// Name https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Name
func (d *Dao) Name() string {
	return d.collection.Name()
}

// Find https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Find
func (d *Dao) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*Cursor, error) {
	return d.collection.Find(ctx, filter, opts...)
}

// FindOne https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.FindOne
func (d *Dao) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *SingleResult {
	return d.collection.FindOne(ctx, filter, opts...)
}

// InsertMany https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany
func (d *Dao) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*InsertManyResult, error) {
	return d.collection.InsertMany(ctx, documents, opts...)
}

// InsertOne https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertOne
func (d *Dao) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*InsertOneResult, error) {
	return d.collection.InsertOne(ctx, document, opts...)
}

// CountDocuments https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.CountDocuments
func (d *Dao) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return d.collection.CountDocuments(ctx, filter, opts...)
}

// DeleteMany https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteMany
func (d *Dao) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*DeleteResult, error) {
	return d.collection.DeleteMany(ctx, filter, opts...)
}

// DeleteOne https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteOne
func (d *Dao) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*DeleteResult, error) {
	return d.collection.DeleteOne(ctx, filter, opts...)
}

// UpdateMany https://github.com/mongodb/mongo-go-driver/blob/master/mongo/collection.go#L548
func (d *Dao) UpdateMany(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*UpdateResult, error) {
	return d.collection.UpdateMany(ctx, filter, update, opts...)
}

// UpdateOne https://github.com/mongodb/mongo-go-driver/blob/master/mongo/collection.go#L532
func (d *Dao) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*UpdateResult, error) {
	return d.collection.UpdateOne(ctx, filter, update, opts...)
}

// ReplaceOne 替换一个
func (d *Dao) ReplaceOne(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (*UpdateResult, error) {
	return d.collection.ReplaceOne(ctx, filter, replacement, opts...)
}

// Distinct  https://github.com/mongodb/mongo-go-driver/blob/master/mongo/collection.go
func (d *Dao) Distinct(ctx context.Context, fieldName string, filter interface{},
	opts ...*options.DistinctOptions) ([]interface{}, error) {
	return d.collection.Distinct(ctx, fieldName, filter, opts...)
}

// FindOneAndUpdate 更新一个
func (d *Dao) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *SingleResult {
	return d.collection.FindOneAndUpdate(ctx, filter, update, opts...)
}

// FindOneAndDelete 删除一个
func (d *Dao) FindOneAndDelete(ctx context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) *SingleResult {
	return d.collection.FindOneAndDelete(ctx, filter, opts...)
}

// UseSession 开启默认事务
func (d *Dao) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return d.collection.Database().Client().UseSession(ctx, fn)
}

// Aggregate 聚合查询
// Deprecated 使用AggregateSafe替换
func (d *Dao) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*Cursor, error) {
	return d.collection.Aggregate(ctx, pipeline, opts...)
}

// BulkWrite 批量写方法
func (d *Dao) BulkWrite(ctx context.Context, models []WriteModel, opts ...*options.BulkWriteOptions) (*BulkWriteResult, error) {
	return d.collection.BulkWrite(ctx, models, opts...)
}

// ----------------------------------------------- 以下为从读模式方法 -----------------------------------------------

// FindSecondaryPreferred https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Find
func (d *Dao) FindSecondaryPreferred(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*Cursor, error) {
	return d.readCollection.Find(ctx, filter, opts...)
}

// FindOneSecondaryPreferred https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.FindOne
func (d *Dao) FindOneSecondaryPreferred(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *SingleResult {
	return d.readCollection.FindOne(ctx, filter, opts...)
}

// CountDocumentsSecondaryPreferred https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.CountDocuments
func (d *Dao) CountDocumentsSecondaryPreferred(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return d.readCollection.CountDocuments(ctx, filter, opts...)
}

// AggregateSecondaryPreferred 聚合查询
func (d *Dao) AggregateSecondaryPreferred(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*Cursor, error) {
	return d.readCollection.Aggregate(ctx, pipeline, opts...)
}
