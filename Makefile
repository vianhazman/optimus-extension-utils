binary_outdir=out
project_name=utils

bin: go build -o ${binary_outdir}/${project_name} .

dist: ./build.sh