/*
 * FairOS-dfs server
 *
 * A list of the currently provided Interfaces to interact with FairOS decentralised file system(dfs), implementing user, pod, file system, key value store and document store
 *
 * API version: v0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiListFileResponse struct {
	Dirs []DirEntry `json:"dirs,omitempty"`
	Files []FileEntry `json:"files,omitempty"`
}
