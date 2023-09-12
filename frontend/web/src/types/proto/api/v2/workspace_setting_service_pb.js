// @generated by protoc-gen-es v1.3.0
// @generated from file api/v2/workspace_setting_service.proto (package slash.api.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message slash.api.v2.WorkspaceSetting
 */
export const WorkspaceSetting = proto3.makeMessageType(
  "slash.api.v2.WorkspaceSetting",
  () => [
    { no: 1, name: "enable_signup", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "resource_relative_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "auto_backup", kind: "message", T: AutoBackupWorkspaceSetting },
  ],
);

/**
 * @generated from message slash.api.v2.AutoBackupWorkspaceSetting
 */
export const AutoBackupWorkspaceSetting = proto3.makeMessageType(
  "slash.api.v2.AutoBackupWorkspaceSetting",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "cron_expression", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "max_keep", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message slash.api.v2.GetWorkspaceSettingRequest
 */
export const GetWorkspaceSettingRequest = proto3.makeMessageType(
  "slash.api.v2.GetWorkspaceSettingRequest",
  [],
);

/**
 * @generated from message slash.api.v2.GetWorkspaceSettingResponse
 */
export const GetWorkspaceSettingResponse = proto3.makeMessageType(
  "slash.api.v2.GetWorkspaceSettingResponse",
  () => [
    { no: 1, name: "setting", kind: "message", T: WorkspaceSetting },
  ],
);

/**
 * @generated from message slash.api.v2.UpdateWorkspaceSettingRequest
 */
export const UpdateWorkspaceSettingRequest = proto3.makeMessageType(
  "slash.api.v2.UpdateWorkspaceSettingRequest",
  () => [
    { no: 1, name: "setting", kind: "message", T: WorkspaceSetting },
    { no: 2, name: "update_mask", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message slash.api.v2.UpdateWorkspaceSettingResponse
 */
export const UpdateWorkspaceSettingResponse = proto3.makeMessageType(
  "slash.api.v2.UpdateWorkspaceSettingResponse",
  () => [
    { no: 1, name: "setting", kind: "message", T: WorkspaceSetting },
  ],
);

