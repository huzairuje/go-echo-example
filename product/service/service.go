package service

import (
	"context"
	"database/sql"
	"github.com/huzairuje/chatat_backend_engineer/product/entity"
	"github.com/huzairuje/chatat_backend_engineer/product/repository"
	"github.com/huzairuje/chatat_backend_engineer/product/request"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"time"
)

type ProductService struct {
	*sql.DB
}

func NewProductService(db *sql.DB) repository.Repository {
	return ProductService{db}
}

const (
	storeQuery    = `INSERT INTO products (title,description,rating,image,created_at) VALUES ($1,$2,$3,$4,$5)`
	listQuery     = `SELECT id,title,description,rating,image,created_at,updated_at,deleted_at FROM products ORDER BY rating DESC`
	findByIdQuery = `SELECT id,title,description,rating,image,created_at,updated_at,deleted_at FROM products WHERE id = $1`
	updateQuery   = `UPDATE products SET title = $1 ,description = $2, rating = $3, image = $4, updated_at = $5 WHERE uid = $6`
	deleteQuery   = `UPDATE products SET delete_at = $1 WHERE uid = $2`
)

func (p ProductService) Store(ctx context.Context, req request.CreateProductRequest) (*entity.Product, error) {
	trx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		logrus.Println("error begin transaction products", err.Error())
		return nil, err
	}
	createdAtNow := time.Now().Format("2006-01-02T15:04:05Z")
	insertId, err := trx.ExecContext(ctx, storeQuery,
		req.Title,
		req.Description,
		req.Rating,
		req.Image,
		createdAtNow)
	if err != nil {
		logrus.Printf("error execute query, error: %s, try rollbacking \n", err.Error())
		err = trx.Rollback()
		if err != nil {
			logrus.Printf("error execute rollback, error: %s ", err.Error())
		}
		return nil, err
	}
	err = trx.Commit()
	if err != nil {
		log.Println("error when commit transaction: ", err.Error())
		return nil, err
	}
	logrus.Println("create product done!")
	idInt64, err := insertId.LastInsertId()
	idString := strconv.FormatInt(idInt64, 10)
	res, err := p.FindByID(idString)
	if err != nil {
		log.Println("error when display object transaction: ", err.Error())
		return nil, err
	}
	return res, nil
}

func (p ProductService) List() ([]*entity.Product, error) {
	var listProducts []*entity.Product
	rows, err := p.DB.Query(listQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		products := new(entity.Product)
		err := rows.Scan(&products.Id,
			&products.Title,
			&products.Description,
			&products.Image,
			&products.Rating,
			&products.CreatedAt,
			&products.UpdatedAt,
			&products.DeletedAt)

		if err != nil {
			log.Println("error scan", err.Error())
		}
		listProducts = append(listProducts, products)
	}
	return listProducts, nil
}

func (p ProductService) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.QueryRow(findByIdQuery, id).Scan(&product.Id,
		&product.Title,
		&product.Description,
		&product.Image,
		&product.Rating,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &product, nil
}

func (p ProductService) Update(ctx context.Context, id string, req request.UpdateProductRequest) (*entity.Product, error) {
	trx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		logrus.Println("error begin transaction products", err.Error())
		return nil, err
	}
	UpdateAtNow := time.Now().Format("2006-01-02T15:04:05Z")
	insertId, err := trx.ExecContext(ctx, updateQuery,
		req.Title,
		req.Description,
		req.Rating,
		req.Image,
		UpdateAtNow,
		id)
	if err != nil {
		logrus.Printf("error execute query, error: %s, try rollbacking \n", err.Error())
		err = trx.Rollback()
		if err != nil {
			logrus.Printf("error execute rollback, error: %s ", err.Error())
		}
		return nil, err
	}
	err = trx.Commit()

	if err != nil {
		log.Println("error when commit transaction: ", err.Error())
		return nil, err
	}
	logrus.Println("update product done!")
	idInt64, err := insertId.RowsAffected()
	idString := strconv.FormatInt(idInt64, 10)
	res, err := p.FindByID(idString)
	if err != nil {
		log.Println("error when display object transaction: ", err.Error())
		return nil, err
	}
	return res, nil
}

func (p ProductService) Destroy(ctx context.Context, id string) error {
	trx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		logrus.Println("error begin transaction products", err.Error())
		return err
	}
	deleteAtNow := time.Now().Format("2006-01-02T15:04:05Z")
	_, err = trx.ExecContext(ctx, deleteQuery, deleteAtNow, id)
	if err != nil {
		logrus.Printf("error execute query, error: %s, try rollbacking \n", err.Error())
		err = trx.Rollback()
		if err != nil {
			logrus.Printf("error execute rollback, error: %s ", err.Error())
		}
		return err
	}
	err = trx.Commit()

	if err != nil {
		log.Println("error when commit transaction: ", err.Error())
		return err
	}
	logrus.Println("delete product done!")
	return nil
}
