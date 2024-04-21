import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";
import { setToastAc } from "../../redux/toastReducer";
import Card from "../Card/Card";
import "./Restaurants.css"

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

class Restaurants extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      
    };
  }


  render() {
      return <div className="restaurant">
        <div className="restaurant_items">
          <Card 
            name="Nacional" 
            distance="08:00-22:00" 
            bonuses="Еда навынос 
            Оплата счета по QR-коду Wi-Fi"
            price="1 800р."
            priceText="Средний чек"
            comments="20 отзывов"
            />
            <Card 
            name="Nacional" 
            distance="08:00-22:00" 
            bonuses="Еда навынос 
            Оплата счета по QR-коду Wi-Fi"
            price="1 800р."
            priceText="Средний чек"
            comments="20 отзывов"
            />
            <Card 
            name="Nacional" 
            distance="08:00-22:00" 
            bonuses="Еда навынос 
            Оплата счета по QR-коду Wi-Fi"
            price="1 800р."
            priceText="Средний чек"
            comments="20 отзывов"
            />
            <Card 
            name="Nacional" 
            distance="08:00-22:00" 
            bonuses="Еда навынос 
            Оплата счета по QR-коду Wi-Fi"
            price="1 800р."
            priceText="Средний чек"
            comments="20 отзывов"
            />
            <Card 
            name="Nacional" 
            distance="08:00-22:00" 
            bonuses="Еда навынос 
            Оплата счета по QR-коду Wi-Fi"
            price="1 800р."
            priceText="Средний чек"
            comments="20 отзывов"
            />
            <Card 
            name="Nacional" 
            distance="08:00-22:00" 
            bonuses="Еда навынос 
            Оплата счета по QR-коду Wi-Fi"
            price="1 800р."
            priceText="Средний чек"
            comments="20 отзывов"
            />
        </div>
      </div>
  }
}

export default connect(mapStateToProps, {setToastAc})(Restaurants);