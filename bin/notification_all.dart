import 'dart:ffi' as ffi;

import '../notify_go.dart' as ng;

void main(List<String> arguments) {
  const _path = './notification_go/notify.a';
  var _n = ng.NativeLibrary(ffi.DynamicLibrary.open(_path));
  _n.notify();
}
