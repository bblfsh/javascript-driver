// @flow

export default class PluginPass {
    testfnc1(a: string, b: ?string, c: number, d: mixed, e: Object){}
    testfnc2(f: string | boolean, g: boolean, h: null, i: void){}
    testfncd(j: "one" | "other"){}

    num: number = 1;
    s: string = "foo";
    pok: File;
}
