package service

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const ENTITYNAME = "Stock"

func GetList(ctx context.Context) (*[]StockEntity, error) {

	var entity []StockEntity
	q := datastore.NewQuery(ENTITYNAME).Order("-TradingValue").Limit(20).Offset(0)
	key, _ := q.GetAll(ctx, &entity)
	if len(key) == 0 {
		return nil, nil
	}
	if err := datastore.GetMulti(ctx, key, entity); err != nil {
		return &entity, err
	}

	return &entity, nil
}

func GetItem(ctx context.Context, code string) (*StockEntity, error) {
	q := datastore.NewQuery(ENTITYNAME).Filter("Code =", code)
	key, _ := q.GetAll(ctx, &[]StockEntity{})
	entity := &StockEntity{}
	if len(key) == 0 {
		return entity, nil
	}
	if err := datastore.Get(ctx, key[0], entity); err != nil {
		return entity, err
	}

	return entity, nil
}

func FullUpdate(ctx context.Context, code string, entity *StockEntity) (*StockEntity, error) {
	q := datastore.NewQuery(ENTITYNAME).Filter("Code =", code)
	key, _ := q.GetAll(ctx, &[]StockEntity{})
	if key == nil {
		key = append(key, datastore.NewIncompleteKey(ctx, ENTITYNAME, nil))
	}
	if _, err := datastore.Put(ctx, key[0], entity); err != nil {
		return &StockEntity{}, err
	}
	return entity, nil
}

func DiffUpdate(ctx context.Context, code string, entity *StockEntity) (*StockEntity, error) {
	q := datastore.NewQuery(ENTITYNAME).Filter("Code =", code)
	key, _ := q.GetAll(ctx, &[]StockEntity{})

	baseEntity := &StockEntity{}
	if err := datastore.Get(ctx, key[0], baseEntity); err != nil {
		return nil, err
	}
	// 差分更新
	baseEntity.Highprice = entity.Highprice
	baseEntity.LowPrice = entity.LowPrice
	baseEntity.ClosingPrice = entity.ClosingPrice
	baseEntity.OpeningPrice = entity.OpeningPrice

	if _, err := datastore.Put(ctx, key[0], baseEntity); err != nil {
		return &StockEntity{}, err
	}
	return baseEntity, nil
}

func DeleteAll(ctx context.Context) (int, error) {
	q := datastore.NewQuery(ENTITYNAME).Limit(1000)
	keys, err := q.GetAll(ctx, &[]StockEntity{})
	if err != nil {
		return 0, err
	}

	delErr := datastore.DeleteMulti(ctx, keys)
	return len(keys), delErr
}
