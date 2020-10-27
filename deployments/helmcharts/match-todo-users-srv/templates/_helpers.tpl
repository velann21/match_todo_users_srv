{{- define "userspods.labels" }}
         labels:
           app: "{{ $.Release.Name }}"
           k8s-app-name: famk8s
           owned_by: platform
{{- end }}


{{- define "usersdeps.labels" }}
   labels:
        app: "{{ $.Release.Name }}"
        k8s-app-name: famk8s
        owned_by: platform
{{- end }}

{{- define "deployment.apiVersion" }}apps/v1{{- end }}

{{- define "service.apiVersion" }}v1{{- end }}

{{- define "service.nameSpace" }}v1{{- end }}