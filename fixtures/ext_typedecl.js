type EnvFunction = {
  (): string,
  <T>((string) => T): T,
  (string): boolean,
  (Array<string>): boolean,
};

const packageToBabelConfig = makeWeakCache(
  (file: ConfigFile): ConfigFile | null => {
    const babel = file.options[("babel": string)];

    if (typeof babel === "undefined") return null;

    if (typeof babel !== "object" || Array.isArray(babel) || babel === null) {
      throw new Error(`${file.filepath}: .babel property must be an object`);
    }

    return {
      filepath: file.filepath,
      dirname: file.dirname,
      options: babel,
    };
  },
);

type VisitorHandler = Function | { enter?: Function, exit?: Function };
export type VisitorMap = {
  [string]: VisitorHandler,
};

export type SimpleCacheConfigurator = SimpleCacheConfiguratorFn &
  SimpleCacheConfiguratorObj;

export function makeWeakCache<
  ArgT: {} | Array<*> | $ReadOnlyArray<*>,
  ResultT,
  SideChannel,
>(
  handler: (ArgT, CacheConfigurator<SideChannel>) => ResultT,
): (ArgT, SideChannel) => ResultT {
  return makeCachedFunction(new WeakMap(), handler);
}

class CacheConfigurator<SideChannel = void> {
  _pairs: Array<[mixed, (SideChannel) => mixed]> = [];
}
