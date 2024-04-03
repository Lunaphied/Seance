{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
  };

  outputs = {self, nixpkgs}: 
  let
    pkgs = nixpkgs.legacyPackages.x86_64-linux;

    transmute = pkgs.callPackage ./package.nix { };
  in {
    devShells.x86_64-linux.default = pkgs.mkShell {
      packages = [
        pkgs.go
      ];
    };

    packages.x86_64-linux = {
      inherit transmute;
      default = transmute;
    };
  };
}
