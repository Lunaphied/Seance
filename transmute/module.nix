{ pkgs, ... }:

let transmute = pkgs.callPackage ./default.nix {};
in {
  services.mattermost = {
    enable = true;
    siteUrl = "http://localhost";
    plugins = [ transmute ];
  };
}
