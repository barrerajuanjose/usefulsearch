package com.finding.usefulsearch.controllers

import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.web.bind.annotation.GetMapping

@Controller
class UsedCarController {

    @GetMapping("/autos-usados-mercadolibre-ultima-oportunidad")
    fun UsedCar(model: Model): String {
        model.addAttribute("title", "PEPE")
        return "used_car";
    }
}