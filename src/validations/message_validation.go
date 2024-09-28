package validations

import (
	"context"
	"github.com/rapradiptya/whatsapp-go/domains/message"
	domainMessage "github.com/rapradiptya/whatsapp-go/domains/message"
	pkgError "github.com/rapradiptya/whatsapp-go/pkg/error"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateRevokeMessage(ctx context.Context, request domainMessage.RevokeRequest) error {
	err := validation.ValidateStructWithContext(ctx, &request,
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.MessageID, validation.Required),
	)

	if err != nil {
		return pkgError.ValidationError(err.Error())
	}

	return nil
}

func ValidateUpdateMessage(ctx context.Context, request domainMessage.UpdateMessageRequest) error {
	err := validation.ValidateStructWithContext(ctx, &request,
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.MessageID, validation.Required),
		validation.Field(&request.Message, validation.Required),
	)

	if err != nil {
		return pkgError.ValidationError(err.Error())
	}

	return nil
}

func ValidateReactMessage(ctx context.Context, request message.ReactionRequest) error {
	err := validation.ValidateStructWithContext(ctx, &request,
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.MessageID, validation.Required),
		validation.Field(&request.Emoji, validation.Required),
	)

	if err != nil {
		return pkgError.ValidationError(err.Error())
	}

	return nil
}

func ValidateDeleteMessage(ctx context.Context, request domainMessage.DeleteRequest) error {
	err := validation.ValidateStructWithContext(ctx, &request,
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.MessageID, validation.Required),
	)

	if err != nil {
		return pkgError.ValidationError(err.Error())
	}

	return nil
}
