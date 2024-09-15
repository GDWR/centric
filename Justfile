build:
    cd website && pnpm run build
    mkdir -p ./static
    cp -r website/dist/* ./static/
    go build
