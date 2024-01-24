FROM golang:1.20-alpine

WORKDIR /src/go_wrapper

COPY --from=jobtype go_wrapper/. /src/go_wrapper/
RUN go get ./... && rm -rf /src/go_wrapper/handler

CMD ./go_wrapper < /dev/null
LABEL racetrack-component="job"


{% for env_key, env_value in env_vars.items() %}
ENV {{ env_key }} "{{ env_value }}"
{% endfor %}

{% if manifest.system_dependencies and manifest.system_dependencies|length > 0 %}
RUN apk add \
    {{ manifest.system_dependencies | join(' ') }}
{% endif %}

{% if manifest.get_jobtype_extra().gomod %}
COPY "{{ manifest.get_jobtype_extra().gomod }}" /src/job/
RUN cd /src/job && go mod download
{% endif %}

COPY . /src/go_wrapper/handler/
RUN chmod -R a+rw /src/go_wrapper && cd /src/go_wrapper/ && go mod download

RUN go get ./... && go build -o go_wrapper

ENV JOB_NAME "{{ manifest.name }}"
ENV JOB_VERSION "{{ manifest.version }}"
ENV GIT_VERSION "{{ git_version }}"
ENV DEPLOYED_BY_RACETRACK_VERSION "{{ deployed_by_racetrack_version }}"
