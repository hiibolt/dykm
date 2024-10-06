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
        clear

        # Custom Colors & Characters
        export RESET='\033[00m'
        export PINK='\033[01;35m'
        export CYAN='\033[01;36m'
        export END='❯'
        export RED='\033[0;31m'

        # Banner Function
        burger() {
          printf "$RESET""[ Nix Flake ]\n"
          printf "$PINK""run$RESET   | ""$PINK""r""$RESET""\n- Starts the Go Webserver\n"
          printf "$PINK""test$RESET  | ""$PINK""t""$RESET""\n- Runs the tests in all files ending in \`test.go\`\n"
          printf "$PINK""build$RESET | ""$PINK""b""$RESET""\n- Builds the Go Webserver, producing \`./dykm\`\n\n"
        }

        # Verify WD is `*/dykm`
        verify_pw() { [[ "$PWD" == */dykm ]] && { return 0; } || { echo "You aren't in the repository root \`./dykm\`!"; return 1; }; }

        # Trap Build Command
        start_application_in_build() {
            # Move into the build directory
            cd ./build || exit

            # Remove any pre-existing Go binary
            rm -f ./dykm

            # Start your application (replace with your actual command)
            go build .; ./dykm & > /dev/null 2>&1

            export SERVER_PID=$!

            # Trap SIGINT (CTRL-C) to handle interruption
            trap 'kill $SERVER_PID; cd ..; printf "\n\nMoving back to root directory...\n\n"' SIGINT

            # Wait for the server to finish
            wait $SERVER_PID > /dev/null 2>&1

            # Optionally, unset the trap after the application exits
            trap - SIGINT
        }

        # Run Command
        # Run Command
        run() {
            verify_pw && start_application_in_build
        }
        alias r="run"

        # Test Command
        alias test="verify_pw && ./build.sh && cd ./build && go test -count=1 . && cd - > /dev/null 2>&1"
        alias t="test"
        
        # Build Command
        alias build="verify_pw && rm -f ./dykm > /dev/null 2>&1 && ./build.sh && cd ./build && go build . && mv ./dykm ../dykm && cd - > /dev/null 2>&1 && ls dykm"
        alias b="build"

        burger

        # Set custom PS1
        export PS1="$RESET$CYAN(nix) $PINK\w$RESET\$(if [ \$? -eq 0 ]; then echo -n '$RESET'; else echo -n '$RED'; fi) $END$RESET "
        '';
      RUST_SRC_PATH = pkgs.rustPlatform.rustLibSrc;
    };
  };
}