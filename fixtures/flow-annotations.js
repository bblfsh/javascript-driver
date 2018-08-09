// @flow

export default class PluginPass {
    testfnc1(a: string, b: ?string, c: number, d: mixed, e: Object){}
    testfnc2(f: string | boolean, g: boolean, h: null, i: void){}
    testfnc3(j: "one" | "other"){}
    testfnc_object(node: ?{
      loc?: { start: { line: number, column: number } },
      _loc?: { start: { line: number, column: number } },
    }) {}

    num: number = 1;
    s: string = "foo";
    pok: File;
    _map: Map<mixed, mixed> = Map();
    
}
