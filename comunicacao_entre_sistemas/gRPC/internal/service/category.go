package service

import (
	"context"
	"io"

    "github.com/guilhermeNatan/full_cycle/14-gRPC/internal/database"
    "github.com/guilhermeNatan/full_cycle/14-gRPC/internal/pb"
)

type CategoryService struct {
    pb.UnimplementedCategoryServiceServer
    CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {

    return &CategoryService{
        CategoryDB: categoryDB,
    }
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description )
    if err != nil {
        return nil, err
    }

    categoryResponse := &pb.Category {
        Id: category.ID,
        Name: category.Name,
        Description: category.Description,
    }

    return categoryResponse, nil
}


func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
    categories , err := c.CategoryDB.FindAll()
    if err != nil {
        return nil, err
    }
    var categoriesResponse []*pb.Category
    for _, category := range categories {
        categoryResponse  := &pb.Category {
            Id:          category.ID,
            Name:        category.Name,
            Description: category.Description,
        }
        categoriesResponse = append(categoriesResponse, categoryResponse)
    }

    return &pb.CategoryList{Categories: categoriesResponse}, nil
}



func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error){
	category, err := c.CategoryDB.Find(in.Id)

    if err != nil {
        return nil, err
    }
    categoryResponse := &pb.Category{
        Id:          category.ID,
        Name:        category.Name,
        Description: category.Description,
    }

    return categoryResponse, nil
}


func (c *CategoryService) CreateCategorySteam(stream pb.CategoryService_CreateCategorySteamServer) error {
   // lista para armazenar os dados no banco
    categories := &pb.CategoryList{}
    // loop infinito para parsistir as categorias
    for {
            category, err := stream.Recv()
            // envia os dados persistidos e fecha a conexão ao final após persistir
    		if err == io.EOF {
    			return stream.SendAndClose(categories)
    		}
    		if err != nil {
    			return err
    		}

    		categoryResult , err := c.CategoryDB.Create(category.Name, category.Description)
            if err != nil {
                return err
            }
          categories.Categories = append(categories.Categories, &pb.Category {
            Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
          })
    }
}

func (c *CategoryService)	CreateCategoryBidirectional(stream pb.CategoryService_CreateCategoryBidirectionalServer) error{
    for {

                category, err := stream.Recv()
                // se chegou no final simplesmente sai do loop
        		if err == io.EOF {
        			return nil
        		}
                if err != nil {
                    return err
                }

                categoryResult , err := c.CategoryDB.Create(category.Name, category.Description)
                if err != nil {
                    return err
                }

                err = stream.Send(&pb.Category{
                            Id:          categoryResult.ID,
                			Name:        categoryResult.Name,
                			Description: categoryResult.Description,
                })
                if err != nil {
                		return err
                }
    }

}
