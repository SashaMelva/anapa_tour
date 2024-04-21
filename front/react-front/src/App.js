import store from './redux/Store';
import React, { useEffect } from 'react';
import { Route, Routes, BrowserRouter, Navigate, NavLink  } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Импорт стилей Bootstrap
import ToastComp from './components/Toast/ToastComp';
import './App.css';
import Registration from './components/Registration/Registration';
import PersonArea from './components/PersonArea/PersonArea'
import Header from './components/Header/Header'
import Footer from './components/Footer/Footer'
import { Provider } from 'react-redux';
import { YMaps } from '@pbe/react-yandex-maps';
import { useSelector } from 'react-redux'; // Импортируем useSelector
import MyMap from './components/MyMap/MyMap';
import Hotels from './components/Hotels/Hotels';
import Restaurants from './components/Restaurants/Restaurants';
import Places from './components/Places/Places';
import Home from './components/Home/Home';
import MainRouteOptions from './components/MainRouteOptions/MainRouteOptions';

function App() {
  return (
    <YMaps query={{
      apikey: "b46407c0-2de2-4342-a545-5a4593b0d161",
      load: "package.full"
    }}>
    <Provider store={store}>
      <InnerApp/>
    </Provider>
    </YMaps>
  );
}

function InnerApp() {
  const isAuth = useSelector(state => state.registration.isAuth); // Получаем значение isAuth из хранилища

  const chests = [
    {
      position: [58.751574, 42.573856],
      box: {},
      org: {}
    },
    {
      position: [52.751574, 41.573856],
      box: {},
      org: {}
    },
    {
      position: [55.751574, 40.573856],
      box: {},
      org: {}
    },
    {
      position: [51.751574, 38.573856],
      box: {},
      org: {}
    }
  ]

  return (
      <div className="wrapper">
        <BrowserRouter>
          <React.Suspense fallback={<div>Загрузка...</div>}>
            <Header />
            <div className="content">
              <Routes>
                <Route path="/" element={<Navigate to={`/registration`} />} />
                <Route path='/registration/' element={<Registration />} />
                <Route path='/personArea' element={isAuth ? <PersonArea /> : <Navigate to={'/registration'}/>} />
                <Route path='/personArea/myMap' element={isAuth ? <MyMap chests={chests}/> : <Navigate to={'/registration'}/>}/>
                <Route path="*" element={<div className='notFound'>404 Страница не найдена</div>} />
                <Route path="/hotels" element={<Hotels/>}/>
                <Route path="/restaurants" element={<Restaurants/>}/>
                <Route path="/places" element={<Places/>}/>
                <Route path="/home" element={<Home/>}/>
                <Route path="/main/routeOptions" element={<MainRouteOptions/>}/>
              </Routes>
            </div>
          </React.Suspense>
          <ToastComp/>
          {/* <Footer/> */}
        </BrowserRouter>
      </div>
  );
}

export default App;