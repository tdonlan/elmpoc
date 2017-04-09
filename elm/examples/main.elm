-- Read more about this program in the official Elm guide:
-- https://guide.elm-lang.org/architecture/effects/http.html

import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Json.Decode as Decode



main =
  Html.program
    { init = init "cats"
    , view = view
    , update = update
    , subscriptions = subscriptions
    }



-- MODEL


type alias Model =
  { topic : String
  , gifUrl : String
  , txt : String
  }


init : String -> (Model, Cmd Msg)
init topic =
  ( Model topic "waiting.gif" "blah some initial text"
  , getRandomGif topic
  )

-- UPDATE


type Msg
  = MorePlease
  | NewGif (Result Http.Error String)
  | PrintText
  | NewText (Result Http.Error String)
    


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    MorePlease ->
      (model, getRandomGif model.topic)

    NewGif (Ok newUrl) ->
      (Model model.topic newUrl model.txt, Cmd.none)

    NewGif (Err _) ->
      (model, Cmd.none)
      
    PrintText ->
      (model, getText)
     
    NewText (Ok newTxt) ->
      (Model model.topic model.gifUrl newTxt, Cmd.none)
      
    NewText (Err _) ->
      (model, Cmd.none)

-- VIEW


view : Model -> Html Msg
view model =
  div []
    [ h2 [] [text model.topic]
    , button [ onClick MorePlease ] [ text "More Please!" ]
    , button [ onClick PrintText ] [ text "Print Text" ]
    , br [] []
    , text model.txt
    , br [] []
    , img [src model.gifUrl] []
    , br [] []
    ]

-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none

-- HTTP

getText : (Cmd Msg)
getText =
  let
    url =
      "api/hello"
  in
    Http.send NewText (Http.getString url)
    
getRandomGif : String -> Cmd Msg
getRandomGif topic =
  let
    url =
      "https://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag=" ++ topic
  in
    Http.send NewGif (Http.get url decodeGifUrl)


decodeGifUrl : Decode.Decoder String
decodeGifUrl =
  Decode.at ["data", "image_url"] Decode.string
