package models

import (
	"fmt"
)

type GameModel struct {
	Game_id     int
	Sub_app_id  int
	Game_name   string
	Create_time string
}

const GAME_MODEL_TABLE_NAME = "game"

func init() {
	fmt.Println("start game_model init")
}

func (g *GameModel) GetGameJson() []GameModel {
	sqlStr := fmt.Sprintf("select * from %s", GAME_MODEL_TABLE_NAME)
	fmt.Println(sqlStr)
	rows, err := db.Query(sqlStr)
	CheckError(err)

	games := make([]GameModel, 0)
	//make([]interface{}, len(columns))

	for rows.Next() {
		var game_id, sub_app_id int
		var game_name, create_time string
		err := rows.Scan(&game_id, &game_name, &create_time, &sub_app_id)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		game_tmp := GameModel{Game_id: game_id, Game_name: game_name, Sub_app_id: sub_app_id, Create_time: create_time}
		fmt.Println(game_tmp)
		games = append(games, game_tmp)
	}

	fmt.Println("end rows")
	return games
}
