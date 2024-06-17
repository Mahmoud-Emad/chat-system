Rails.application.routes.draw do
  resources :applications, param: :token, only: [ :index, :show, :create, :update ] do
    resources :chats, param: :number, only: [ :index, :show, :create ] do
      resources :messages, param: :number, only: [ :index, :show, :create ] do
        get "search", on: :collection
      end
    end
  end
end
