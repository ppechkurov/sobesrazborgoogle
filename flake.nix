{
  description = "Sobes flake development environment";

  # Flake inputs
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  # Flake outputs
  outputs =
    {
      self,
      nixpkgs,
      utils,
    }:
    utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      with pkgs;
      {
        # Development environment output
        devShells = {
          default = mkShell {
            # The Nix packages provided in the environment
            packages = [
              go-task
              go_1_27
              gofumpt
              golangci-lint
              golangci-lint-langserver
              gopls
              gotools
              pgcli
              sqlc
              sqlite
              watchexec
            ];
          };
        };

        packages.default = buildGo127Module {
          pname = "api";
          version = "0.0.0";
          src = ./.;
          vendorHash = "sha256-uPqabZgQGQulf+F3BvMLhv4O0h5jOq12F7K60u5xjtA=";
          ldflags = [
            "-s"
            "-w"
          ];
        };
      }
    );
}
