FROM moby/buildkit:v0.9.3
WORKDIR /bank
COPY bank README.md /bank/
ENV PATH=/bank:$PATH
ENTRYPOINT [ "/bhojpur/bank" ]