// This is just a stub for IDE.
// It gets replaced by user's Job code in wrappers/docker/golang/job-template.Dockerfile
package stub

import (
	log "github.com/inconshreveable/log15"
)

func Perform(input map[string]interface{}) (interface{}, error) {
	log.Info("I wish only to serve...", log.Ctx{"input": input})
	return nil, nil
}
