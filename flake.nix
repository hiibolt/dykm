{
  description = "Development environment for Python";
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
      shellHook = "";
      RUST_SRC_PATH = pkgs.rustPlatform.rustLibSrc;
    };
  };
}