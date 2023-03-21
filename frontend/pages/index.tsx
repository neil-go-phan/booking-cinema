import type { NextPage } from 'next';
import { useEffect, useContext } from 'react';
import { checkLogin } from '@/helpers/checkLogin';
import { IsLoggedContext } from '@/helpers/isLoggedContext';

const Home: NextPage = () => {
  const logged = useContext(IsLoggedContext);
  useEffect(() => {
    async function isLogged() {
      if(await checkLogin()){
        return logged.setIsLogged(true);
      }
      return logged.setIsLogged(false);
    }
    isLogged()
  }, [logged]);
  return (
    <div>
      This is home page localhost:3000/ 
      <p>isLogged: {`${logged.isLogged}`}</p>
    </div>
  )
}

export default Home;