import { useEffect, useState } from 'react';

const useRequest = (
  api: () => Promise<any>,
  autoLoad: boolean | undefined = false,
): {
  readonly data: any;
  load: () => Promise<any>;
  readonly loading: boolean;
} => {
  let [data, setData] = useState({});
  let [loading, setLoading] = useState(true);
  const load = () => {
    setLoading(true);
    return api()
      .then(res => {
        if (res) {
          setData(res.data.data);
        }
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    if (autoLoad) {
      load();
    }
  }, []);

  return { data, load, loading };
};
export default useRequest;
