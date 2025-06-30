{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
    in {
      packages = {
        default = pkgs.buildGoModule {
          inherit version;
          pname = "steam-exporter";
          src = ./.;
          vendorHash = "sha256-KBHAkFkvv5BwAyn/qqCF6RMBRzyN+oi1/vECIDPA4tI=";
        };
      };

      devShells = {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            cobra-cli
          ];
        };
      };
    });
}
