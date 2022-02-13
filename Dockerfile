# SPDX-License-Identifier: Apache-2.0
# SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
# SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
# SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#    http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.17

WORKDIR /build
ADD . /build/

RUN mkdir /tmp/cache
RUN CGO_ENABLED=0 GOCACHE=/tmp/cache go build -mod=mod -v -o /tmp/mks-server .

FROM scratch

LABEL org.opencontainers.image.authors="Avinal Kumar <avinkuma@redhat.com>, Utkarsh Chaurasia <uchauras@redhat.com>, Satyam Bhardwaj <sabhardw@redhat.com>"
LABEL org.opencontainers.image.source="https://github.com/MiniTeks/mks-server"
LABEL org.opencontainers.image.version="latest"

WORKDIR /app
COPY --from=0 /tmp/mks-server /app/mks-server
COPY ./config /app/config

ENTRYPOINT ["/app/mks-server"]