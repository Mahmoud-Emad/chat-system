class ApplicationsController < ActionController::API
  def index
    applications = Application.cached_all
    render json: applications.as_json(only: [:name, :token, :created_at, :updated_at])
  end

  def show
    application = Application.find_by(token: params[:token])
    if application
      render json: application.as_json(
        only: [ :name, :token, :created_at, :updated_at ],
        include: {
          chats: {
            only: [ :number, :created_at, :updated_at, :messages_count ]
          }
        }
      )
    else
      render json: { error: "Application is not found" }, status: :not_found
    end
  end

  def create
    if params["name"].blank?
      render json: { error: "Application name is required, Try to send `{ name: 'Application Name!' }` as the request body." }, status: 400
      return
    end

    application = Application.new(application_params)

    if application.save
      render json: application.as_json(
        only: [ :name, :token, :created_at, :updated_at ],
        include: {
          chats: {
            only: [ :number, :created_at, :updated_at, :messages_count ]
          }
        }
      ), status: :created
    else
      render json: application.errors, status: :unprocessable_entity
    end
  end

  def update
    application = Application.find_by(token: params[:token])
    if application
      if application.update(application_params)
        render json: application.as_json(only: [ :name, :token, :created_at, :updated_at ]), status: :created
      else
        render json: application.errors, status: :unprocessable_entity
      end
    else
      render json: { error: "Application is not found" }, status: :not_found
    end
  end

  private
  def application_params
    params.require(:application).permit(:name)
  end
end
