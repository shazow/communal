{
  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem
    (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
    in {
      devShells.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go
          pkgs.golint
        ];

        shellHook = ''
          export PS1="[dev] $PS1"

          [[ -f .env ]] && source .env
        '';
      };
    });
}
