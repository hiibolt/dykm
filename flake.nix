{
  description = "Development environment for Go";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };
  outputs = { self, nixpkgs }: 
    let
      system = "x86_64-linux";
      
      pkgs = import nixpkgs { 
        inherit system;
      };
    in
    {
    devShells.x86_64-linux.default = pkgs.mkShell {
      nativeBuildInputs = with pkgs; [ openssl.dev pkg-config ];
      buildInputs = with pkgs; [ go cargo rustc rustfmt rust-analyzer clippy ];
      shellHook = ''
        alias run="cd ./backend/src/dykm; go run .; cd -"
        alias test="cd ./backend/src/dykm; go test -count=1 .; cd -"
        '';
      RUST_SRC_PATH = pkgs.rustPlatform.rustLibSrc;
    };
  };
}