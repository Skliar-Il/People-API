package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Skliar-Il/People-API/internal/dto"
	"github.com/Skliar-Il/People-API/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
	"go.uber.org/zap"
)

type PeopleClientInterface interface {
	GetAge(c context.Context, name string) (int, error)
	GetGender(ctx context.Context, name string) (string, error)
	GetNationalize(ctx context.Context, name string) (string, error)
}

type PeopleClientLinkConfig struct {
	AgeLink         string `env:"LINK_AGE"`
	NationalizeLink string `env:"LINK_NATIONALIZE"`
	GenderLink      string `env:"LINK_GENDER"`
}
type PeopleClient struct {
	Link   *PeopleClientLinkConfig
	Client *client.Client
}

func NewPeopleClient(client_ *client.Client, linkConfig *PeopleClientLinkConfig) *PeopleClient {
	return &PeopleClient{
		Client: client_,
		Link:   linkConfig,
	}
}

func (p *PeopleClient) GetAge(ctx context.Context, name string) (int, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)

	res, err := p.Client.Get(p.Link.AgeLink + fmt.Sprintf("?name=%s", name))
	if err != nil {
		localLogger.Error(ctx, "get age error", zap.Error(err))
		return 0, fmt.Errorf("get age error")
	}
	if res == nil {
		localLogger.Error(ctx, "failed get age, nil body")
		return 0, nil
	}
	if res.StatusCode() != fiber.StatusOK {
		localLogger.Error(ctx, "unexpected status code",
			zap.Int("status", res.StatusCode()),
			zap.String("body", string(res.Body())),
			zap.Error(err))
		return 0, fmt.Errorf("unexpected status code age")
	}

	var body dto.PeopleAgeClientDTO
	if err := json.Unmarshal(res.Body(), &body); err != nil {
		localLogger.Error(ctx, "unmarshal age body error",
			zap.String("body", string(res.Body())),
			zap.Error(err))
		return 0, fmt.Errorf("unmarshal body error")
	}

	return body.Age, nil
}

func (p *PeopleClient) GetGender(ctx context.Context, name string) (string, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)

	res, err := p.Client.Get(p.Link.GenderLink + fmt.Sprintf("?name=%s", name))
	if err != nil {
		localLogger.Error(ctx, "get gender client error", zap.Error(err))
	}
	if res == nil {
		localLogger.Error(ctx, "failed get gender, nil body")
		return "", nil
	}
	if res.StatusCode() != fiber.StatusOK {
		localLogger.Error(ctx, "unexpected status code gender",
			zap.Int("status", res.StatusCode()),
			zap.String("body", string(res.Body())))
		return "", fmt.Errorf("unexpected status code")
	}

	var body dto.PeopleGenderClientDTO
	if err := json.Unmarshal(res.Body(), &body); err != nil {
		localLogger.Error(ctx, "unmarshal gender body error",
			zap.String("body", string(res.Body())),
			zap.Error(err))
		return "", fmt.Errorf("unmarshal body error")
	}

	if body.Gender == nil {
		return "", nil
	}

	return *body.Gender, nil

}

func (p *PeopleClient) GetNationalize(ctx context.Context, name string) (string, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)

	res, err := p.Client.Get(p.Link.NationalizeLink + fmt.Sprintf("?name=%s", name))
	if err != nil {
		localLogger.Error(ctx, "get nationalize error", zap.Error(err))
		return "", fmt.Errorf("get nationalize error")
	}
	if res == nil {
		localLogger.Error(ctx, "failed get nationalize, nil body")
		return "", nil
	}
	if res.StatusCode() != fiber.StatusOK {
		localLogger.Error(ctx, "unexpected status code nationalize",
			zap.Int("status", res.StatusCode()),
			zap.String("body", string(res.Body())))
		return "", fmt.Errorf("unexpected status code")
	}

	var body dto.PeopleNationalizeClientDTO
	if err := json.Unmarshal(res.Body(), &body); err != nil {
		localLogger.Error(ctx, "unmarshal nationalize body error",
			zap.String("body", string(res.Body())),
			zap.Error(err))
		return "", fmt.Errorf("unmarshal body error")
	}

	if len(body.Country) == 0 {
		return "", nil
	}

	return body.Country[0].CountryId, nil
}
