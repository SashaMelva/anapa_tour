import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import "./Card.css"

const mapStateToProps = (state) => {
  
}

const Card = (props) => {
  return <div className="card_itemBlock">
      <div className="card_item">
        <div className="card_header">
        <div className="card_like"/>
        </div>
        <div className="card_body">
          <div className="card_info">
            <div className="card_left">
              <div className="card_name">{props.name}</div>
              <div className="card_dist">
                <div className="card_distImg"/>
                <div className="card_distText">{props.distance}</div>
              </div>
              <div className="card_bonuses">{props.bonuses} </div>
            </div>
            <div className="card_right">
              <div className="card_rating">
                <div className="card_starsImg"/>
                <div className="card_comments">{props.comments}</div>
              </div>
              <div className="card_priceText">{props.priceText}: {props.price}</div>
            </div>
          </div>
          <Button className="card_bonusBtn">Бонусная программа</Button>
        </div>
      </div>
  </div>
}

export default connect(mapStateToProps)(Card);