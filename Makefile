gen:
	openapi-generator generate  --api-name-suffix ledgera -g go --additional-properties=prependFormOrBodyParameters=true -o ./ -i api/openapi.yaml