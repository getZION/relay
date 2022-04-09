package validator

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/getzion/relay/api"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func Test_ShouldValidate_Community(t *testing.T) {
	InitValidator()

	faker.SetGenerateUniqueValues(false)

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "Name field should be required",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.Name = ""
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Description field should be maximum 250 character",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				faker.SetRandomStringLength(300)
				faker.FakeData(&model.Description)
				faker.SetRandomStringLength(25)
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "OwnerDid field should be required",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.OwnerDid = ""
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "OwnerUsername field should be required",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.OwnerUsername = ""
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "EscrowAmount field should be greater or equal to zero",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.EscrowAmount = -1
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "EscrowAmount field should be less than 100000",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.EscrowAmount = 100000
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PricePerMessage field should be greater or equal to zero",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PricePerMessage = -1
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PricePerMessage field should be less than 100000",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PricePerMessage = 100000
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PriceToJoin field should be greater or equal to zero",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PriceToJoin = -1
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PriceToJoin field should be less than 100000",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PriceToJoin = 100000
				model.Tags = []string{}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Tags field should be less than 5",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.Tags = []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6"}
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				model := api.Community{
					Name:            "test",
					Description:     "test",
					EscrowAmount:    10,
					OwnerDid:        "test_did",
					OwnerUsername:   "test_username",
					PricePerMessage: 2,
					PriceToJoin:     5,
				}
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := Struct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				validationError := err.(validator.ValidationErrors)
				require.Len(t, validationError, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ShouldValidate_JoinCommunity(t *testing.T) {
	InitValidator()

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "userdid and community zid cannot be empty",
			generate: func() interface{} {
				return api.JoinCommunity{
					UserDid:      "",
					CommunityZid: "",
				}
			},
			expectedError:      true,
			expectedErrorCount: 2,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				return api.JoinCommunity{
					UserDid:      "test_did",
					CommunityZid: "test_zid",
				}
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := Struct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				validationError := err.(validator.ValidationErrors)
				require.Len(t, validationError, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ShouldValidate_LeaveCommunity(t *testing.T) {
	InitValidator()

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "userdid and community zid cannot be empty",
			generate: func() interface{} {
				return api.LeaveCommunity{
					UserDid:      "",
					CommunityZid: "",
				}
			},
			expectedError:      true,
			expectedErrorCount: 2,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				return api.LeaveCommunity{
					UserDid:      "test_did",
					CommunityZid: "test_zid",
				}
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := Struct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				validationError := err.(validator.ValidationErrors)
				require.Len(t, validationError, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ShouldValidate_Conversation(t *testing.T) {
	InitValidator()

	faker.SetGenerateUniqueValues(false)

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "CommunityZid field should be required",
			generate: func() interface{} {
				model := api.Conversation{}
				faker.FakeData(&model)
				model.CommunityZid = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "UserDid field should be required",
			generate: func() interface{} {
				model := api.Conversation{}
				faker.FakeData(&model)
				model.UserDid = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Text or Link fields should be required",
			generate: func() interface{} {
				model := api.Conversation{}
				faker.FakeData(&model)
				model.Text = ""
				model.Link = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 2,
		},
		{
			name: "should be valid with only text",
			generate: func() interface{} {
				model := api.Conversation{}
				faker.FakeData(&model)
				model.Link = ""
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
		{
			name: "should be valid with only link",
			generate: func() interface{} {
				model := api.Conversation{}
				faker.FakeData(&model)
				model.Text = ""
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				model := api.Conversation{}
				faker.FakeData(&model)
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := Struct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				validationError := err.(validator.ValidationErrors)
				require.Len(t, validationError, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ShouldValidate_User(t *testing.T) {
	InitValidator()

	faker.SetGenerateUniqueValues(false)

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "Name field should be required",
			generate: func() interface{} {
				model := api.User{}
				faker.FakeData(&model)
				model.Email = faker.Email()
				model.Username = faker.Username()
				model.Name = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Email field should be invalid",
			generate: func() interface{} {
				model := api.User{}
				faker.FakeData(&model)
				model.Username = faker.Username()
				model.Email = "test-"
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Username field should be at least 6 character",
			generate: func() interface{} {
				model := api.User{}
				faker.FakeData(&model)
				model.Username = "usern"
				model.Email = faker.Email()
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Username field should be maximum 16 character",
			generate: func() interface{} {
				model := api.User{}
				faker.FakeData(&model)
				model.Username = "maximumsixteencharacter"
				model.Email = faker.Email()
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Username field should be alphanumeric and underscore",
			generate: func() interface{} {
				model := api.User{}
				faker.FakeData(&model)
				model.Username = "test*user!"
				model.Email = faker.Email()
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				model := api.User{}
				faker.FakeData(&model)
				model.Username = "test_user"
				model.Email = faker.Email()
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
		{
			name: "Email can be empty",
			generate: func() interface{} {
				model := api.User{}
				faker.FakeData(&model)
				model.Username = faker.Username()
				model.Email = ""
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := Struct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				validationError := err.(validator.ValidationErrors)
				require.Len(t, validationError, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ShouldValidate_Comment(t *testing.T) {
	InitValidator()

	tests := []struct {
		name               string
		model              api.Comment
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name:               "conversation should be required",
			model:              api.Comment{UserDid: "did1", Text: "test"},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name:               "userdid should be required",
			model:              api.Comment{ConversationZid: "zid1", Text: "test"},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name:               "text or link should be required",
			model:              api.Comment{ConversationZid: "zid1", UserDid: "did1"},
			expectedError:      true,
			expectedErrorCount: 2,
		},
		{
			name:               "should be valid with only text",
			model:              api.Comment{ConversationZid: "zid1", UserDid: "did1", Text: "test"},
			expectedError:      false,
			expectedErrorCount: 0,
		},
		{
			name:               "should be valid with only link",
			model:              api.Comment{ConversationZid: "zid1", UserDid: "did1", Link: "test"},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := Struct(tt.model)

			if tt.expectedError {
				require.NotNil(t, err)
				validationError := err.(validator.ValidationErrors)
				require.Len(t, validationError, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
