{
  buildGoModule,
  lib
}:

buildGoModule {
  pname = "seance-transmute";
  version = "0.0.1";

  src = lib.cleanSource ./.;

  vendorHash = "sha256-K/C2vNj92zkF7gGeSWarru0lJsJVt1+ZmkhzlocecJ0=";

  meta = with lib; {
    description = "A ritual to channel the unseen";
    homepage = "https://github.com/Qyriad/Seance";
    license = licenses.mit;
  };

  installPhase = ''
    runHook preInstall
    mkdir server
    mv $GOPATH/bin/transmute server
    tar -cJf $out server plugin.json
    runHook postInstall
  '';
}
