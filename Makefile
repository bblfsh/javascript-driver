-include .sdk/Makefile

$(if $(filter true,$(sdkloaded)),,$(error You must install bblfsh-sdk))

test-native-internal:
	cd native; \
	yarn && yarn test

build-native-internal:
	cd native; \
	yarn && yarn build && \
	cp lib/index.js $(BUILD_PATH)/index.js && \
	cp native.sh $(BUILD_PATH)/native && \
	cp -R node_modules $(BUILD_PATH)/node_modules && \
	chmod +x $(BUILD_PATH)/native
