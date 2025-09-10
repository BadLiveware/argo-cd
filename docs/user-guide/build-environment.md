# Build Environment

[Custom tools](../operator-manual/config-management-plugins.md), [Helm](helm.md), [Jsonnet](jsonnet.md), and [Kustomize](kustomize.md) support the following build env vars:

| Variable                            | Description                                                               |
| ----------------------------------- | ------------------------------------------------------------------------- |
| `ARGOCD_APP_NAME`                   | The name of the application.                                              |
| `ARGOCD_APP_NAMESPACE`              | The destination namespace of the application.                             |
| `ARGOCD_APP_PROJECT_NAME`           | The name of the project the application belongs to.                       |
| `ARGOCD_APP_REVISION`               | The resolved revision, e.g. `f913b6cbf58aa5ae5ca1f8a2b149477aebcbd9d8`.   |
| `ARGOCD_APP_REVISION_SHORT`         | The resolved short revision, e.g. `f913b6c`.                              |
| `ARGOCD_APP_REVISION_SHORT_8`       | The resolved short revision with length 8, e.g. `f913b6cb`.               |
| `ARGOCD_APP_SOURCE_PATH`            | The path of the app within the source repo.                               |
| `ARGOCD_APP_SOURCE_REPO_URL`        | The source repo URL.                                                      |
| `ARGOCD_APP_SOURCE_TARGET_REVISION` | The target revision from the spec, e.g. `master`.                         |
| `ARGOCD_APP_ORIGINAL_REVISION`      | The pre-resolved targeted revision, e.g. `v1.0.*`, or `HEAD`, or `master` |
| `KUBE_VERSION`                      | The semantic version of Kubernetes without trailing metadata.             |
| `KUBE_API_VERSIONS`                 | The version of the Kubernetes API.                                        |

The follow build env vars are available in certain situations, e.g. targetRevision requires a [range resolution](../user-guide/tracking_strategies)
| Variable | Description | `RESOLUTION_TYPE`
| ---------------------------- | ----------------------------------------------------------------------- |----------------------|
| `ARGOCD_APP_RESOLUTION_TYPE` | The type of resolution done for `revisionTarget`, e.g. `tag`, or `truncated_commit_sha` | Is `RESOLUTION_TYPE` |
| `ARGOCD_APP_RESOLVED_TAG` | The tag resolved for [range resolution](../user-guide/tracking_strategies) | `tag` |
| `ARGOCD_APP_RESOLVED_TO` | The post-resolved target revision, e.g. `v1.0.1` | `symbolic_reference` |

In case you don't want a variable to be interpolated, `$` can be escaped via `$$`.

```
command:
  - sh
  - -c
  - |
    echo $$FOO
```
