package lepora_test

import (
	"testing"

	"github.com/Augani/lepora"
	"github.com/stretchr/testify/assert"
)

//test lepora
func TestLepora(t *testing.T) {
	//create a new lepora instance
	lep, err := lepora.Setup(lepora.LeporaOptions{
		Method: lepora.Local,
		Name: "AppTest",
		MaxSize: 1024,
		MaxFiles: 5,
		MaxDays: 7,
		Debug: true,
	})
	assert.NoError(t, err)
	assert.NotNil(t, lep)
	//log a message
	lep.Log("Message", "Hellow world")

	//check for a file created in the current directory with the string in log
	

}