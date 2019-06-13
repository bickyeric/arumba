package telegraph_test

import (
	"errors"
	"testing"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/mocks"
	"github.com/bickyeric/arumba/service/telegraph"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewCreatePage(t *testing.T) {
	creator := telegraph.NewCreatePage()
	assert.NotNil(t, creator)
}

func TestPerform_ErrorPosting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedNetwork := mocks.NewMockNetworkInterface(ctrl)
	mockedNetwork.EXPECT().POST(gomock.Any(), gomock.Any()).Return(nil, errors.New("ini error"))
	connection.TelegraphNetwork = mockedNetwork

	source := "Mangacan"
	title := "One Piece 1 | Episode-nya"
	images := []string{
		"https://www.mangatail.me/sites/default/files/manga/5/284497//20190612153018763.jpg",
		"https://www.mangatail.me/sites/default/files/manga/5/284497//20190612153018763.jpg",
		"https://www.mangatail.me/sites/default/files/manga/5/284497//20190612153018763.jpg",
		"https://www.mangatail.me/sites/default/files/manga/5/284497//20190612153018763.jpg",
	}

	cp := telegraph.NewCreatePage()
	_, err := cp.Perform(source, title, images)
	assert.NotNil(t, err)
}

func TestPerform_ErrorUnmarshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedNetwork := mocks.NewMockNetworkInterface(ctrl)
	mockedNetwork.EXPECT().POST(gomock.Any(), gomock.Any()).Return([]byte("invalid_json"), nil)
	connection.TelegraphNetwork = mockedNetwork

	source := "Mangacan"
	title := "One Piece 1 | Episode-nya"
	images := []string{}

	cp := telegraph.NewCreatePage()
	_, err := cp.Perform(source, title, images)
	assert.NotNil(t, err)
}

func TestPerform(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedNetwork := mocks.NewMockNetworkInterface(ctrl)
	mockedNetwork.EXPECT().POST(gomock.Any(), gomock.Any()).Return([]byte("{}"), nil)
	connection.TelegraphNetwork = mockedNetwork

	source := "Mangacan"
	title := "One Piece 1 | Episode-nya"
	images := []string{}

	cp := telegraph.NewCreatePage()
	_, err := cp.Perform(source, title, images)
	assert.Nil(t, err)
}
