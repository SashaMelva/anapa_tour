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

import { useSelector } from 'react-redux'; // Импортируем useSelector
import MyMap from './components/MyMap/MyMap';

function App() {
  return (
    <Provider store={store}>
      <InnerApp/>
    </Provider>
  );
}

function InnerApp() {
  const isAuth = useSelector(state => state.registration.isAuth); // Получаем значение isAuth из хранилища

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
                <Route path='/personArea/myMap' element={isAuth ? <MyMap /> : <Navigate to={'/registration'}/>} />
                <Route path="*" element={<div className='notFound'>Спят тюбики</div>} />
              </Routes>
            </div>
          </React.Suspense>
          <ToastComp/>
          <Footer/>
        </BrowserRouter>
      </div>
  );
}

export default App;