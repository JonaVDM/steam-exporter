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
    })
    // {
      nixosModules.default = {
        config,
        pkgs,
        lib,
        ...
      }:
        with lib; let
          cfg = config.jo.services.steam-exporter;
        in {
          options.jo.services.steam-exporter = {
            enable = mkEnableOption "Enable the Steam exporter";
            host = mkOption {
              type = types.string;
              default = ":6718";
            };

            envFile = mkOption {
              type = types.path;
            };

            userid = mkOption {
              type = types.string;
            };
          };

          config = mkIf cfg.enable {
            systemd.services.steam-exporter = {
              description = "Steam Exporter";
              wantedBy = ["multi-user.target"];
              environment = {
                STEAM_EXPORTER_HOST = cfg.host;
                STEAM_EXPORTER_USER = cfg.userid;
              };

              serviceConfig = {
                DynamicUser = "yes";
                ExecStart = "${self.packages.${pkgs.system}.default}/bin/steam-exporter";
                Restart = "on-failure";
                RestartSec = "5s";
                EnvironmentFile = cfg.envFile;
              };
            };
          };
        };
    };
}
