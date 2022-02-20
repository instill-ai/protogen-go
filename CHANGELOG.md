# Changelog

## [0.1.0-alpha](https://github.com/instill-ai/protogen-go/compare/v0.0.0-alpha...v0.1.0-alpha) (2022-02-19)


### Features

* add conversion ([6b27ee3](https://github.com/instill-ai/protogen-go/commit/6b27ee33fca9c8cd0115fa8b38b4befcff946ea4))
* add DeleteModel method ([#14](https://github.com/instill-ai/protogen-go/issues/14)) ([9a2713d](https://github.com/instill-ai/protogen-go/commit/9a2713d704b2c30945696afd7c5289e63b80ee64))
* add liveness and readiness for pipeline ([767f166](https://github.com/instill-ai/protogen-go/commit/767f166cb546e69115d178855092b0a8186e3200))
* add rpc service of triggering pipeline by uploading file ([4b21d58](https://github.com/instill-ai/protogen-go/commit/4b21d582913832af0f3df694590621f5b78dd306))
* initialize repo ([0d6d613](https://github.com/instill-ai/protogen-go/commit/0d6d613615394a388150ed8610306f20b9599309))
* **model-backend:** add generated code for `model-backend` ([#10](https://github.com/instill-ai/protogen-go/issues/10)) ([2203592](https://github.com/instill-ai/protogen-go/commit/22035920e4553e4897436de5d6d17d6932830209))
* **model:** add generate protobuf files for model-backend ([#5](https://github.com/instill-ai/protogen-go/issues/5)) ([a4cd01f](https://github.com/instill-ai/protogen-go/commit/a4cd01f3b2f952d5256db11e251de07a2cd0b6e8))


### Bug Fixes

* change go module name ([f56fdee](https://github.com/instill-ai/protogen-go/commit/f56fdeee9a34b7acb0949db238ea9ec6b1b6b891))
* change id to uint64 ([5eaf34d](https://github.com/instill-ai/protogen-go/commit/5eaf34df589c6df096f667dbb4ac8c0612e82cf0))
* change message content and fix optional issue ([b74ec60](https://github.com/instill-ai/protogen-go/commit/b74ec602a8d26b6bf9e20f4ece7b45bc2fdff2b1))
* correct protobuffer sequence number of triggering pipeline request ([a0bfca2](https://github.com/instill-ai/protogen-go/commit/a0bfca2fd3c95a19a25789f5ff5fa9ce4d1f031b))
* model version status into enum ([#13](https://github.com/instill-ai/protogen-go/issues/13)) ([23f8fbf](https://github.com/instill-ai/protogen-go/commit/23f8fbfa8e13acaae9b3d6ceee01fbcb71f259ae))
* move cv task from inference to create model method ([#15](https://github.com/instill-ai/protogen-go/issues/15)) ([36f2370](https://github.com/instill-ai/protogen-go/commit/36f2370961ef89fbc127b9d9505aa8bb245e57d4))
* **pipeline-backend:** remove `TriggerPipeline` endpoint ([#7](https://github.com/instill-ai/protogen-go/issues/7)) ([9ef454b](https://github.com/instill-ai/protogen-go/commit/9ef454b9b8089d48b6816a19668462ce6ae79c36))
* **pipeline-backend:** use Any instead of Struct ([#6](https://github.com/instill-ai/protogen-go/issues/6)) ([1216e1c](https://github.com/instill-ai/protogen-go/commit/1216e1c0631dad691105f316354b497817e6ef20))
* **pipeline:** change `Any` back to `Struct` ([#9](https://github.com/instill-ai/protogen-go/issues/9)) ([09ba6de](https://github.com/instill-ai/protogen-go/commit/09ba6de03b2719294998b167acaffe3dbe18c6ca))
* **pipeline:** change `model_name` to `name` ([#4](https://github.com/instill-ai/protogen-go/issues/4)) ([e6c93ae](https://github.com/instill-ai/protogen-go/commit/e6c93ae50935fa50bb893f92c351ad5c9ae83d5d))
* **pipeline:** remove unused code ([#28](https://github.com/instill-ai/protogen-go/issues/28)) ([3c0c45e](https://github.com/instill-ai/protogen-go/commit/3c0c45e118585703f0bacb154b0d2b1bcd78223b))
* remove unuse code ([9f931a0](https://github.com/instill-ai/protogen-go/commit/9f931a0df89739a23533352b3ff4ab7786e81c11))
* remove version from creating model, the version will be auto increased ([#25](https://github.com/instill-ai/protogen-go/issues/25)) ([6a4b524](https://github.com/instill-ai/protogen-go/commit/6a4b524722dcf8be9eec6bdade197adef79cedc5))
* use `full_name` instead of `fullName` ([#26](https://github.com/instill-ai/protogen-go/issues/26)) ([5b8ade9](https://github.com/instill-ai/protogen-go/commit/5b8ade9e492384edd7467f5840ca8e4b5af4fc7d))
* use consistent response of list pipeline ([70bcb11](https://github.com/instill-ai/protogen-go/commit/70bcb112cf4e73179d528bba72c43c695432b889))
* use FieldMask as updating flag ([be17192](https://github.com/instill-ai/protogen-go/commit/be1719222c9d8ec6e84b43495bebacf179ae10d1))
* use name as query parameter instead of id ([b0e9531](https://github.com/instill-ai/protogen-go/commit/b0e9531885c2aed7e9a0d5d9c22e45558b8104e6))
* use original timestamppb ([78653a9](https://github.com/instill-ai/protogen-go/commit/78653a95256dbbc47069cda74f639bc3438e17ba))
* wrong health path ([#12](https://github.com/instill-ai/protogen-go/issues/12)) ([9b9b707](https://github.com/instill-ai/protogen-go/commit/9b9b7079641441e2f6d0bdf90400a27b289959c9))
