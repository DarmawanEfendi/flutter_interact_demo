package main

import (
	"fmt"
	"gihub.com/DarmawanEfendi/flutter_interact_demo/models"
	"gihub.com/DarmawanEfendi/flutter_interact_demo/restapi"
	"gihub.com/DarmawanEfendi/flutter_interact_demo/restapi/operations"
	"gihub.com/DarmawanEfendi/flutter_interact_demo/restapi/operations/profile"
	"github.com/bxcodec/faker/v3"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

type Profile struct {
	ID          int64
	Name        *string `faker:"name"`
	Url         *string `faker:"url"`
	IsFavourite *bool
}

func main() {
	data := make([]Profile, 0)
	for i := 1; i <= 100; i += 1 {
		prof := &Profile{}
		err := faker.FakeData(prof)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, *prof)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSwaggerProfileAPI(swaggerSpec)
	api.ProfileGetProfilesHandler = profile.GetProfilesHandlerFunc(func(params profile.GetProfilesParams) middleware.Responder {
		profileData := make([]*models.ProfileSchema, 0)
		for _, prof := range data {
			flag := true
			if params.OnlyFavourite != nil {
				if *params.OnlyFavourite != *prof.IsFavourite {
					flag = false
				}
			}

			if flag == true {
				model := &models.ProfileSchema{
					ID:          prof.ID,
					Name:        prof.Name,
					URL:         prof.Url,
					IsFavourite: prof.IsFavourite,
				}
				profileData = append(profileData, model)
			}
		}

		message := "Successfully get list of profiles"
		status := int64(200)
		return profile.NewGetProfilesOK().WithPayload(&models.ProfilesResponse{
			Data: profileData,
			Meta: &models.MetaSchema{
				Message: &message,
				Status:  &status,
			},
		})
	})

	api.ProfilePostProfileFavouriteHandler = profile.PostProfileFavouriteHandlerFunc(func(params profile.PostProfileFavouriteParams) middleware.Responder {
		index := -1
		for i := 0; i < len(data); i += 1 {
			if params.Body.ID == data[i].ID {
				data[i].IsFavourite = params.Body.IsFavourite
				index = i
				break
			}
		}

		if index == -1 {
			message := "Profile not found"
			status := int64(404)
			return profile.NewPostProfileFavouriteNotFound().WithPayload(&models.GeneralErrorResponse{
				Meta: &models.MetaSchema{
					Message: &message,
					Status:  &status,
				},
			})
		}


		message := "Successfully changes saved"
		status := int64(201)
		return profile.NewPostProfileFavouriteCreated().WithPayload(&models.ProfileResponse{
			Data: &models.ProfileSchema{
				ID:          data[index].ID,
				IsFavourite: data[index].IsFavourite,
				Name:        data[index].Name,
				URL:         data[index].Url,
			},
			Meta: &models.MetaSchema{
				Message: &message,
				Status:  &status,
			},
		})
	})
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	defer server.Shutdown()

	server.Port = 8080
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
