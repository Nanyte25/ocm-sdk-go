/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalGitlabIdentityProvider writes a value of the 'gitlab_identity_provider' type to the given writer.
func MarshalGitlabIdentityProvider(object *GitlabIdentityProvider, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeGitlabIdentityProvider(object, stream)
	stream.Flush()
	return stream.Error
}

// writeGitlabIdentityProvider writes a value of the 'gitlab_identity_provider' type to the given stream.
func writeGitlabIdentityProvider(object *GitlabIdentityProvider, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.ca != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("ca")
		stream.WriteString(*object.ca)
		count++
	}
	if object.url != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("url")
		stream.WriteString(*object.url)
		count++
	}
	if object.clientID != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("client_id")
		stream.WriteString(*object.clientID)
		count++
	}
	if object.clientSecret != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("client_secret")
		stream.WriteString(*object.clientSecret)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalGitlabIdentityProvider reads a value of the 'gitlab_identity_provider' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalGitlabIdentityProvider(source interface{}) (object *GitlabIdentityProvider, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readGitlabIdentityProvider(iterator)
	err = iterator.Error
	return
}

// readGitlabIdentityProvider reads a value of the 'gitlab_identity_provider' type from the given iterator.
func readGitlabIdentityProvider(iterator *jsoniter.Iterator) *GitlabIdentityProvider {
	object := &GitlabIdentityProvider{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "ca":
			value := iterator.ReadString()
			object.ca = &value
		case "url":
			value := iterator.ReadString()
			object.url = &value
		case "client_id":
			value := iterator.ReadString()
			object.clientID = &value
		case "client_secret":
			value := iterator.ReadString()
			object.clientSecret = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
