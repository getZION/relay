package collections

import (
	"testing"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/dwn"
	"github.com/getzion/relay/api/dwn/handler"
	"github.com/getzion/relay/api/models"
	"github.com/getzion/relay/api/storage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_User_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	storage.EXPECT().GetUsers().Times(1).Return([]api.User{
		{Id: 1, Did: "did1"},
		{Id: 2, Did: "did2"},
	}, nil)
	modelManager := models.NewModelManager(storage)

	tests := []struct {
		name                 string
		message              *dwn.Message
		expectedErrorMessage string
	}{
		{
			name: "should return 2 user",
			message: &dwn.Message{
				Descriptor: &dwn.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Method:   constants.COLLECTIONS_QUERY,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{ModelManager: modelManager, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)
		})
	}
}
