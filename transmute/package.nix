{
  lib,
  buildGoModule,
  runCommandNoCC
}:

let module = buildGoModule {
  pname = "seance-transmute";
  version = "0.0.1";

  src = lib.cleanSource ./.;

  vendorHash = "sha256-JKgMkx+QTt2PzA4kJi5zBDJHtPt8pHs+UtasenRzbT4=";
};
  #meta = with lib; {
  #  description = "A ritual to channel the unseen";
  #  homepage = "https://github.com/Qyriad/Seance";
  #  license = licenses.mit;
  #};
in runCommandNoCC "seance-transmute.tgz" {} ''
  mkdir seance-transmute
  cat ${./plugin.json} > seance-transmute/plugin.json
  cp ${module}/bin/seance-transmute seance-transmute/seance-transmute
  chmod ug+rw seance-transmute/seance-transmute
  tar -czvf $out seance-transmute
''
