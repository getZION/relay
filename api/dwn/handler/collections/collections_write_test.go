package collections

import (
	"testing"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/dwn"
	"github.com/getzion/relay/api/dwn/handler"
	"github.com/getzion/relay/api/models"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_UserCreate(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	modelManager := models.NewModelManager(storage)

	storage.EXPECT().InsertUser(gomock.Any()).Times(1).Return(nil)

	entries, err := CollectionsWrite(&handler.RequestContext{
		ModelManager: modelManager,
		Message: &dwn.Message{
			Data: `{ "Name": "test_name", "Username": "test_username", "Did": "did", "Model": "Zion.User.V1" }`,
			Descriptor: &dwn.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Len(t, entries, 1)
}
