{
  description = "OnePdfPlease - A simple tui for working with pdfs";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        pname = "onepdfplease";
        version = "0.1.0";
      in
      {
        packages.default = pkgs.buildGoModule {
          inherit pname version;
          src = ./.;

          vendorHash = "sha256-NIo8O1BsQP0J80QBjqB6M/2ihFu9R+MIpZR6MU1jO2c=";

          proxyVendor = true;

          ldflags = [
            "-s"
            "-w"
            "-X main.version=${version}"
            "-X main.commit=${self.rev or "dirty"}"
          ];

          meta = with pkgs.lib; {
            description = "A tui for working with pdfs";
            license = licenses.mit;
            mainProgram = pname;
            platforms = platforms.linux ++ platforms.darwin;
          };
        };

      }
    );
}
